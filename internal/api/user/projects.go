// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package user

import (
	"net/http"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/api/request"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"
)

// HandleProjects returns an http.HandlerFunc that writes a json-encoded
// list of projects to the response body.
func HandleProjects(projects store.ProjectStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		viewer, _ := request.UserFrom(r.Context())
		list, err := projects.List(r.Context(), viewer.ID, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("email", viewer.Email).
				Errorf("cannot list projects")
		} else {
			render.JSON(w, list, 200)
		}
	}
}
