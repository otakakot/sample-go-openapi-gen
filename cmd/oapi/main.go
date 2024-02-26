package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echomiddleware "github.com/oapi-codegen/echo-middleware"

	"github.com/otakakot/sample-go-openapi-gen/pkg/oapi"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "1323"
	}

	e := echo.New()

	e.Use(middleware.Logger())

	e.Use(middleware.Recover())

	swagger, err := oapi.GetSwagger()
	if err != nil {
		panic(err)
	}

	// おまじない ref: https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/echo/petstore.go#L30-L32
	swagger.Servers = nil

	options := &echomiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(ctx context.Context, ai *openapi3filter.AuthenticationInput) error {
				slog.Info(fmt.Sprintf("authentication func: %+v", ai))

				switch ai.SecuritySchemeName {
				case "bearerAuth":
					slog.Info(fmt.Sprintf("bearer token: %s", ai.RequestValidationInput.Request.Header.Get("Authorization")))

					authorization := ai.RequestValidationInput.Request.Header.Get("Authorization")

					authorizations := strings.Split(authorization, " ")

					if len(authorizations) != 2 {
						return fmt.Errorf("invalid authorization: %s", authorization)
					}

					if authorizations[0] != "Bearer" {
						return fmt.Errorf("invalid token: %s", authorization)
					}

					if authorizations[1] != "token" {
						return fmt.Errorf("invalid token: %s", authorizations[1])
					}
				case "cookieAuth":
					cookie, err := ai.RequestValidationInput.Request.Cookie("SESSION")
					if err != nil {
						return fmt.Errorf("cookie not found: %w", err)
					}

					slog.Info(fmt.Sprintf("cookie: %s", cookie.Value))
				default:
					return fmt.Errorf("unknown security scheme: %s", ai.SecuritySchemeName)
				}

				return nil
			},
		},
	}

	e.Use(echomiddleware.OapiRequestValidatorWithOptions(swagger, options))

	oapi.RegisterHandlers(e, &Server{
		pets: map[int64]oapi.Pet{},
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		slog.Info("start server listen")

		if err := e.Start(":" + port); err != nil && errors.Is(err, http.ErrServerClosed) {
			e.Logger.Error("shutting down the server")
		}
	}()

	<-ctx.Done()

	slog.Info("start server shutdown")

	ctx, cansel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cansel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Panic(err)
	}

	slog.Info("done server shutdown")
}

var _ oapi.ServerInterface = (*Server)(nil)

type Server struct {
	pets map[int64]oapi.Pet
}

// CreatePets implements oapi.ServerInterface.
func (srv *Server) CreatePets(ctx echo.Context) error {
	var pet oapi.Pet

	if err := ctx.Bind(&pet); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	slog.Info(fmt.Sprintf("create pet: %v", pet))

	srv.pets[pet.Id] = pet

	return ctx.JSON(http.StatusCreated, pet)
}

// DeleteSession implements oapi.ServerInterface.
func (*Server) DeleteSession(ctx echo.Context) error {
	ck := http.Cookie{
		Name:   "SESSION",
		Value:  "",
		MaxAge: -1,
	}

	ctx.SetCookie(&ck)

	return ctx.JSON(http.StatusOK, nil)
}

// GetSession implements oapi.ServerInterface.
func (*Server) GetSession(ctx echo.Context) error {
	ck := http.Cookie{
		Name:  "SESSION",
		Value: uuid.NewString(),
	}

	ctx.SetCookie(&ck)

	return ctx.JSON(http.StatusOK, nil)
}

// ListPets implements oapi.ServerInterface.
func (srv *Server) ListPets(ctx echo.Context, params oapi.ListPetsParams) error {
	slog.Info(fmt.Sprintf("list pets: %d", *params.Limit))

	pets := make([]oapi.Pet, 0, *params.Limit)

	for _, pet := range srv.pets {
		pets = append(pets, pet)
	}

	ctx.Response().Header().Set("X-Next", "next")

	return ctx.JSON(http.StatusOK, pets)
}

// Redirect implements oapi.ServerInterface.
func (*Server) Redirect(ctx echo.Context) error {
	return ctx.Redirect(http.StatusFound, "https://example.com")
}

// ShowPetById implements oapi.ServerInterface.
func (srv *Server) ShowPetById(ctx echo.Context, petId string) error {
	petID, err := strconv.ParseInt(petId, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	pet, ok := srv.pets[petID]
	if !ok {
		return ctx.JSON(http.StatusNotFound, nil)
	}

	return ctx.JSON(http.StatusOK, pet)
}
