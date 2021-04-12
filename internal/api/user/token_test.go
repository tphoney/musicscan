// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/api/request"
	"github.com/tphoney/musicscan/types"
	"github.com/dgrijalva/jwt-go"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestToken(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
		Token: "12345",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleToken(nil)(w, r)
	if got, want := w.Code, 200; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	result := &types.Token{}
	json.NewDecoder(w.Body).Decode(&result)

	_, err := jwt.Parse(result.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(mockUser.Token), nil
	})
	if err != nil {
		t.Error(err)
	}
}

// the purpose of this unit test is to verify that an error
// updating the database will result in an internal server
// error returned to the client.
func TestToken_UpdateError(t *testing.T) {
	generateOriginal := generate
	generate = func(sub int64, secret string) (string, error) {
		return "", jwt.ErrInvalidKey
	}
	defer func() {
		generate = generateOriginal
	}()

	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleToken(nil)(w, r)
	if got, want := w.Code, 500; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := new(render.Error), &render.Error{Message: "Failed to generate token"}
	json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) != 0 {
		t.Errorf(diff)
	}
}
