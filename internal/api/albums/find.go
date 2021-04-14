// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package albums

import (
	"net/http"
	"strconv"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"

	"github.com/go-chi/chi"
)

// HandleFind returns an http.HandlerFunc that writes the
// json-encoded artist details to the response body.
func HandleFind(artists store.ArtistStore, albums store.AlbumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectID, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		artistID, err := strconv.ParseInt(chi.URLParam(r, "artist"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse artist id")
			return
		}

		albumID, err := strconv.ParseInt(chi.URLParam(r, "album"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse album id")
			return
		}

		artist, err := artists.Find(r.Context(), artistID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", artistID).
				Debugln("artist not found")
			return
		}

		album, err := albums.Find(r.Context(), albumID)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", albumID).
				Debugln("artist not found")
			return
		}

		if artist.Project != projectID {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("artist", artistID).
				WithField("album", albumID).
				WithField("project", projectID).
				Debugln("project id mismatch")
			return
		}

		if artist.ID != album.Artist {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("artist.id", artist.ID).
				WithField("album.id", album.ID).
				WithField("project", projectID).
				Debugln("artist id mismatch")
			return
		}

		render.JSON(w, album, 200)
	}
}
