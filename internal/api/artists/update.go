// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artists

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/go-chi/chi"
)

// HandleUpdate returns an http.HandlerFunc that processes http
// requests to update the object details.
func HandleUpdate(artists store.artistStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		project, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		id, err := strconv.ParseInt(chi.URLParam(r, "artist"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse artist id")
			return
		}

		in := new(types.artistInput)
		err = json.NewDecoder(r.Body).Decode(in)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("project", project).
				WithField("id", id).
				Debugln("cannot unmarshal json request")
			return
		}

		artist, err := artists.Find(r.Context(), id)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", id).
				Debugln("artist not found")
			return
		}

		if artist.Project != project {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("id", id).
				WithField("project", project).
				Debugln("project id mismatch")
			return
		}

		if in.Name.IsZero() == false {
			artist.Name = in.Name.String
		}
		if in.Desc.IsZero() == false {
			artist.Desc = in.Desc.String
		}

		err = artists.Update(r.Context(), artist)
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("name", artist.Name).
				WithField("id", id).
				Errorln("cannot update the artist")
		} else {
			render.JSON(w, artist, 200)
		}
	}
}
