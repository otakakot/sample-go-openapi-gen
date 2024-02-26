package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/google/uuid"
	"github.com/ogen-go/ogen/middleware"

	"github.com/otakakot/sample-go-openapi-gen/pkg/ogen"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	hdl, err := ogen.NewServer(
		&Handler{
			pets: map[int64]ogen.Pet{},
		},
		&SecurityHandler{},
		ogen.WithMiddleware(AcccessLog()),
	)
	if err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           hdl,
		ReadHeaderTimeout: 30 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	go func() {
		slog.Info("start server listen")

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()

	slog.Info("start server shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		panic(err)
	}

	slog.Info("done server shutdown")
}

var _ ogen.Handler = (*Handler)(nil)

type Handler struct {
	pets map[int64]ogen.Pet
}

// CreatePets implements ogen.Handler.
func (hdl *Handler) CreatePets(ctx context.Context, req *ogen.Pet) (ogen.CreatePetsRes, error) {
	slog.Info(fmt.Sprintf("create pet: %+v", req))

	hdl.pets[req.ID] = *req

	return &ogen.CreatePetsCreated{}, nil
}

// DeleteSession implements ogen.Handler.
func (*Handler) DeleteSession(ctx context.Context) (*ogen.DeleteSessionOK, error) {
	ck := http.Cookie{
		Name:   "SESSION",
		Value:  "",
		MaxAge: -1,
	}

	return &ogen.DeleteSessionOK{
		SetCookie: ogen.NewOptString(ck.String()),
	}, nil
}

// GetSession implements ogen.Handler.
func (*Handler) GetSession(ctx context.Context) (*ogen.GetSessionOK, error) {
	ck := http.Cookie{
		Name:  "SESSION",
		Value: uuid.NewString(),
	}

	return &ogen.GetSessionOK{
		SetCookie: ogen.NewOptString(ck.String()),
	}, nil
}

// ListPets implements ogen.Handler.
func (hdl *Handler) ListPets(ctx context.Context, params ogen.ListPetsParams) (*ogen.PetsHeaders, error) {
	slog.Info(fmt.Sprintf("limit: %d", params.Limit.Value))

	response := make(ogen.Pets, 0, params.Limit.Value)

	for _, pet := range hdl.pets {
		response = append(response, pet)
	}

	return &ogen.PetsHeaders{
		XNext:    ogen.NewOptString("next"),
		Response: response,
	}, nil
}

// NewError implements ogen.Handler.
func (*Handler) NewError(ctx context.Context, err error) *ogen.ErrorStatusCode {
	return &ogen.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: ogen.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		},
	}
}

// Redirect implements ogen.Handler.
func (*Handler) Redirect(ctx context.Context) (ogen.RedirectRes, error) {
	uri, _ := url.Parse("http://example.com")

	return &ogen.RedirectFound{
		Location: ogen.NewOptURI(*uri),
	}, nil
}

// ShowPetById implements ogen.Handler.
func (hdl *Handler) ShowPetById(ctx context.Context, params ogen.ShowPetByIdParams) (ogen.ShowPetByIdRes, error) {
	petID, err := strconv.ParseInt(params.PetID, 10, 64)
	if err != nil {
		return &ogen.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}, nil
	}

	pet, ok := hdl.pets[petID]
	if !ok {
		return &ogen.Error{
			Code:    http.StatusNotFound,
			Message: "not found",
		}, nil
	}

	return &ogen.Pet{
		ID:   pet.ID,
		Name: pet.Name,
		Tag:  pet.Tag,
	}, nil
}

var _ ogen.SecurityHandler = (*SecurityHandler)(nil)

type SecurityHandler struct{}

// HandleBearerAuth implements ogen.SecurityHandler.
func (*SecurityHandler) HandleBearerAuth(ctx context.Context, operationName string, t ogen.BearerAuth) (context.Context, error) {
	if t.Token == "" {
		return ctx, errors.New("token is empty")
	}

	return ctx, nil
}

// HandleCookieAuth implements ogen.SecurityHandler.
func (*SecurityHandler) HandleCookieAuth(ctx context.Context, operationName string, t ogen.CookieAuth) (context.Context, error) {
	if t.APIKey == "" {
		return ctx, errors.New("api key is empty")
	}

	return ctx, nil
}

func AcccessLog() ogen.Middleware {
	return func(req middleware.Request, next middleware.Next) (middleware.Response, error) {
		slog.Info(fmt.Sprintf("%s %s", req.Raw.Method, req.Raw.URL))

		res, err := next(req)
		if err != nil {
			slog.Error(fmt.Sprintf("Error %v", err))
		} else {
			slog.Info(fmt.Sprintf("Response %T", res.Type))
		}

		return res, err
	}
}
