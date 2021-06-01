// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package projects

import (
	"net/http"
	"strconv"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"

	"github.com/go-chi/chi"
)

func HandleFind(projects store.ProjectStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		project, err := projects.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("project not found")
		} else {
			render.JSON(w, project, 200)
		}
	}
}

func HandleFindBadAlbums(projects store.ProjectStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		project, err := projects.FindBadAlbums(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("project not found")
		} else {
			render.JSON(w, project, 200)
		}
	}
}
