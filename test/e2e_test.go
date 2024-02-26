package test_test

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"

	"github.com/otakakot/sample-go-openapi-gen/pkg/ogen"
)

func TestE2E(t *testing.T) {
	t.Parallel()

	endpoint := os.Getenv("ENDPOINT")
	if endpoint == "" {
		endpoint = "http://localhost:1324"
	}

	cli, err := ogen.NewClient(endpoint, &SecuritySource{})
	if err != nil {
		t.Fatalf("Error creating client: %v\n", err)
	}

	ctx := context.Background()

	t.Run("Pets", func(t *testing.T) {
		t.Parallel()

		want := ogen.Pet{
			ID:   1,
			Name: "name",
			Tag:  ogen.NewOptString("tag"),
		}

		if _, err := cli.CreatePets(ctx, &want); err != nil {
			t.Errorf("Error posting pet: %v\n", err)
		}

		res, err := cli.ShowPetById(ctx, ogen.ShowPetByIdParams{
			PetID: fmt.Sprintf("%d", 1),
		})

		got, ok := res.(*ogen.Pet)
		if !ok {
			t.Errorf("Error showing pet by id: %v\n", err)

			return
		}

		if diff := cmp.Diff(want, *got); diff != "" {
			t.Errorf("(-want, +got)\n%s\n", diff)
		}

		list, err := cli.ListPets(ctx, ogen.ListPetsParams{
			Limit: ogen.NewOptInt32(10),
		})
		if err != nil {
			t.Errorf("Error listing pets: %v\n", err)
		}

		if list.XNext.Value != "next" {
			t.Errorf("Expected next, got %s\n", list.XNext.Value)
		}

		if len(list.Response) != 1 {
			t.Errorf("Expected 1 pet, got %d\n", len(list.Response))
		}

		if diff := cmp.Diff(want, list.Response[0]); diff != "" {
			t.Errorf("(-want, +got)\n%s\n", diff)
		}

		if _, err := cli.ListPets(ctx, ogen.ListPetsParams{
			Limit: ogen.NewOptInt32(101),
		}); err == nil {
			t.Errorf("Expected error listing pets, got nil\n")
		}
	})

	t.Run("Session", func(t *testing.T) {
		t.Parallel()

		res, err := cli.GetSession(ctx)
		if err != nil {
			t.Errorf("Error getting session: %v\n", err)
		}

		if res.GetSetCookie().Value == "" {
			t.Errorf("Expected non-empty session cookie, got empty\n")
		}

		// http.Client が内部で勝手にリダイレクト処理して型が合わないのでエラーとなる
		// if _, err := cli.Redirect(ctx); err != nil {
		// 	t.Errorf("Error redirecting: %v\n", err)
		// }

		rqst, err := http.NewRequest(http.MethodGet, endpoint+"/redirect", nil)
		if err != nil {
			t.Fatal(err)
		}

		rqst.AddCookie(&http.Cookie{
			Name:  "SESSION",
			Value: "session", // とりあえず適当な値
		})

		resp, err := http.DefaultClient.Do(rqst)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("unexpected status code: %d", resp.StatusCode)
		}

		if _, err := cli.DeleteSession(ctx); err != nil {
			t.Errorf("Error deleting session: %v\n", err)
		}
	})
}

var _ ogen.SecuritySource = (*SecuritySource)(nil)

type SecuritySource struct{}

// BearerAuth implements ogen.SecuritySource.
func (*SecuritySource) BearerAuth(ctx context.Context, operationName string) (ogen.BearerAuth, error) {
	return ogen.BearerAuth{
		Token: "token",
	}, nil
}

// CookieAuth implements ogen.SecuritySource.
func (*SecuritySource) CookieAuth(ctx context.Context, operationName string) (ogen.CookieAuth, error) {
	return ogen.CookieAuth{
		APIKey: uuid.NewString(),
	}, nil
}
