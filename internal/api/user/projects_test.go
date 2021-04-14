// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/api/request"
	"github.com/tphoney/musicscan/mocks"
	"github.com/tphoney/musicscan/types"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestProjectList(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
	}

	mockList := []*types.Project{
		{
			Name: "test",
			Desc: "desc",
		},
	}

	projs := mocks.NewMockProjectStore(controller)
	projs.EXPECT().List(gomock.Any(), mockUser.ID, gomock.Any()).Return(mockList, nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleProjects(projs)(w, r)
	if got, want := w.Code, http.StatusOK; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := []*types.Project{}, mockList
	_ = json.NewDecoder(w.Body).Decode(&got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}

func TestProjectListErr(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockUser := &types.User{
		ID:    1,
		Email: "octocat@github.com",
	}

	projs := mocks.NewMockProjectStore(controller)
	projs.EXPECT().List(gomock.Any(), mockUser.ID, gomock.Any()).Return(nil, render.ErrNotFound)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	r = r.WithContext(
		request.WithUser(r.Context(), mockUser),
	)

	HandleProjects(projs)(w, r)
	if got, want := w.Code, http.StatusInternalServerError; want != got {
		t.Errorf("Want response code %d, got %d", want, got)
	}

	got, want := &render.Error{}, render.ErrNotFound
	_ = json.NewDecoder(w.Body).Decode(got)
	if diff := cmp.Diff(got, want); len(diff) > 0 {
		t.Errorf(diff)
	}
}
