// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreatePets implements createPets operation.
//
// Create pet.
//
// POST /pets
func (UnimplementedHandler) CreatePets(ctx context.Context, req *Pet) (r CreatePetsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteSession implements deleteSession operation.
//
// Delete session.
//
// DELETE /session
func (UnimplementedHandler) DeleteSession(ctx context.Context) (r *DeleteSessionOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetSession implements getSession operation.
//
// Get session.
//
// GET /session
func (UnimplementedHandler) GetSession(ctx context.Context) (r *GetSessionOK, _ error) {
	return r, ht.ErrNotImplemented
}

// ListPets implements listPets operation.
//
// List pets.
//
// GET /pets
func (UnimplementedHandler) ListPets(ctx context.Context, params ListPetsParams) (r *PetsHeaders, _ error) {
	return r, ht.ErrNotImplemented
}

// Redirect implements redirect operation.
//
// Redirect.
//
// GET /redirect
func (UnimplementedHandler) Redirect(ctx context.Context) (r RedirectRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ShowPetById implements showPetById operation.
//
// Get pet by id.
//
// GET /pets/{pet_id}
func (UnimplementedHandler) ShowPetById(ctx context.Context, params ShowPetByIdParams) (r ShowPetByIdRes, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
