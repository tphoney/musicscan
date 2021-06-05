// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package projects

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/go-chi/chi"
)

// HandleFind returns an http.HandlerFunc that writes the
// json-encoded project details to the response body.
func HandleScan(artistStore store.ArtistStore, albumStore store.AlbumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projID, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		basePath, err := ioutil.ReadDir("/media/tp/stuff/Music")
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("id", projID).
				Debugln("unable to read directory")
		}
		newArtists := 0
		newAlbums := 0
		changedAlbums := 0
		for _, f := range basePath {
			if !f.IsDir() {
				continue
			}

			var foundArtist *types.Artist
			artistPath := fmt.Sprintf("/media/tp/stuff/Music/%s", f.Name())
			// try to find the artist
			foundArtist, artistFindErr := artistStore.FindByName(r.Context(), f.Name())
			if artistFindErr != nil {
				// artist not found create it
				inArtist := &types.Artist{
					Name:    f.Name(),
					Project: projID,
					Desc:    artistPath,
					Wanted:  true,
				}
				artistCreateErr := artistStore.Create(r.Context(), inArtist)
				if artistCreateErr != nil {
					render.InternalError(w, artistCreateErr)
					logger.FromRequest(r).
						WithError(artistCreateErr).
						WithField("id", projID).
						Debugln("unable to create artist")
				}
				var err2 error
				foundArtist, err2 = artistStore.FindByName(r.Context(), f.Name())
				if err2 != nil {
					render.InternalError(w, err2)
					logger.FromRequest(r).
						WithError(err2).
						WithField("id", projID).
						Debugln("unable to find artist after creation")
				}
				newArtists++
			}
			albumPaths, _ := ioutil.ReadDir(artistPath)
			for _, albumPath := range albumPaths {
				if albumPath.IsDir() {
					dbAlbum, findAlbumErr := albumStore.FindByName(r.Context(), foundArtist.ID, albumPath.Name())
					if findAlbumErr != nil {
						// new album
						abs := artistPath + "/" + albumPath.Name()
						mp3Matches, _ := filepath.Glob(abs + "/*.mp3")
						flacMatches, _ := filepath.Glob(abs + "/*.flac")
						format := ""
						if len(mp3Matches) != 0 && len(flacMatches) != 0 {
							format = "mp3+flac"
						} else if len(mp3Matches) != 0 {
							format = "mp3"
						} else if len(flacMatches) != 0 {
							format = "flac"
						}
						inputAlbum := &types.Album{
							Name:   albumPath.Name(),
							Artist: foundArtist.ID,
							Desc:   abs,
							Format: format,
							Wanted: false,
						}
						createAlbumErr := albumStore.Create(r.Context(), inputAlbum)
						if createAlbumErr != nil {
							render.InternalError(w, createAlbumErr)
							logger.FromRequest(r).
								WithError(createAlbumErr).
								WithField("id", projID).
								Debugln("unable to create album")
						}
						newAlbums++
					} else {
						// lets see if anything has changed
						abs := artistPath + "/" + albumPath.Name()
						mp3Matches, _ := filepath.Glob(abs + "/*.mp3")
						flacMatches, _ := filepath.Glob(abs + "/*.flac")
						format := ""
						if len(mp3Matches) != 0 && len(flacMatches) != 0 {
							format = "mp3+flac"
						} else if len(mp3Matches) != 0 {
							format = "mp3"
						} else if len(flacMatches) != 0 {
							format = "flac"
						}
						change := false
						if dbAlbum.Format != format {
							change = true
						}
						if change {
							dbAlbum.Format = format
							dbAlbum.Wanted = false
							updateAlbumErr := albumStore.Update(r.Context(), dbAlbum)
							if updateAlbumErr != nil {
								render.InternalError(w, updateAlbumErr)
								logger.FromRequest(r).
									WithError(updateAlbumErr).
									WithField("id", projID).
									Debugln("unable to update album")
							}
							changedAlbums++
						}
					}
				}
			}
		}
		outputMessage := fmt.Sprintf("New artists %d, New Albums %d", newArtists, newAlbums)
		render.JSON(w, outputMessage, 200)
	}
}
