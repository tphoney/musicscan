// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artists

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/tphoney/musicscan/internal/api/render"
	"github.com/tphoney/musicscan/internal/logger"
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/go-chi/chi"
)

type ArtistResponse struct {
	Artists struct {
		Href  string `json:"href"`
		Items []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Followers struct {
				Href  interface{} `json:"href"`
				Total int64       `json:"total"`
			} `json:"followers"`
			Genres []string `json:"genres"`
			Href   string   `json:"href"`
			ID     string   `json:"id"`
			Images []struct {
				Height int64  `json:"height"`
				URL    string `json:"url"`
				Width  int64  `json:"width"`
			} `json:"images"`
			Name       string `json:"name"`
			Popularity int64  `json:"popularity"`
			Type       string `json:"type"`
			URI        string `json:"uri"`
		} `json:"items"`
		Limit    int64       `json:"limit"`
		Next     string      `json:"next"`
		Offset   int64       `json:"offset"`
		Previous interface{} `json:"previous"`
		Total    int64       `json:"total"`
	} `json:"artists"`
}

type AlbumsResponse struct {
	Href  string `json:"href"`
	Items []struct {
		AlbumGroup string `json:"album_group"`
		AlbumType  string `json:"album_type"`
		Artists    []struct {
			ExternalUrls struct {
				Spotify string `json:"spotify"`
			} `json:"external_urls"`
			Href string `json:"href"`
			ID   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
			URI  string `json:"uri"`
		} `json:"artists"`
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		ID     string `json:"id"`
		Images []struct {
			Height int64  `json:"height"`
			URL    string `json:"url"`
			Width  int64  `json:"width"`
		} `json:"images"`
		Name                 string `json:"name"`
		ReleaseDate          string `json:"release_date"`
		ReleaseDatePrecision string `json:"release_date_precision"`
		TotalTracks          int64  `json:"total_tracks"`
		Type                 string `json:"type"`
		URI                  string `json:"uri"`
	} `json:"items"`
	Limit    int64       `json:"limit"`
	Next     interface{} `json:"next"`
	Offset   int64       `json:"offset"`
	Previous interface{} `json:"previous"`
	Total    int64       `json:"total"`
}

type WebAlbums struct {
	Items []struct {
		Name    string
		Year    string
		Spotify string
	}
}

var oauth string = "BQCy9YkTOglsg4N1sYPInzcqKWHrgkvyD6fYyEjz9dNMfuph1A1Oe76v90BOndp5w4dt60GJ2YTjPSD71YQZb_OXOvboVz6pPb2WeceJ_ijX0d9g0sFn3vlXznwTIwZr-feG1PyDBXxcx1z-94bJnuSEJvNZTgE"

func spotifyLookupArtist(artistName, oauth string) (spotifyID string, err error) {
	urlEncodedArtist := url.QueryEscape(artistName)
	artistURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist&limit=1", urlEncodedArtist)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, artistURL, nil)
	bearer := fmt.Sprintf("Bearer %s", oauth)
	req.Header.Add("Authorization", bearer)
	response, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("lookupArtist: get failed from spotify: %s", err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("lookupArtist: unable to parse response from spotify: %s", err.Error())
	}
	//fmt.Println(string(body))
	var artistResponse ArtistResponse
	_ = json.Unmarshal(body, &artistResponse)
	if len(artistResponse.Artists.Items) == 0 {
		return "", fmt.Errorf("no artists match '%s', on spotify", artistName)
	}
	//fmt.Println(artistResponse.Artists.Items[0].ID)
	return artistResponse.Artists.Items[0].ID, nil
}

func spotifyLookupArtistAlbums(artistSpotify, oauth string) (albums WebAlbums, err error) {
	albumURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums?include_groups=album&market=us&limit=50&", artistSpotify)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, albumURL, nil)
	bearer := fmt.Sprintf("Bearer %s", oauth)
	req.Header.Add("Authorization", bearer)
	response, err := client.Do(req)
	if err != nil {
		return albums, fmt.Errorf("lookupArtistAlbums: get failed from spotify: %s", err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return albums, fmt.Errorf("lookupArtistAlbums: unable to parse response from spotify: %s", err.Error())
	}
	var albumsResponse AlbumsResponse
	_ = json.Unmarshal(body, &albumsResponse)

	for i := range albumsResponse.Items {
		year := albumsResponse.Items[i].ReleaseDate[0:4]
		albums.Items = append(albums.Items, struct {
			Name    string
			Year    string
			Spotify string
		}{Name: albumsResponse.Items[i].Name, Year: year, Spotify: albumsResponse.Items[i].ID})
	}

	return albums, nil
}

func stringMatcher(dbName, webName string) (matched bool) {
	// first attempt at a match check everything lower case
	cleanedDB := strings.ToLower(dbName)
	cleanedWeb := strings.ToLower(webName)
	cleanedDB = strings.TrimSpace(cleanedDB)
	cleanedWeb = strings.TrimSpace(cleanedWeb)
	if cleanedDB == cleanedWeb {
		return true
	}
	// remove the dbname from the album name, trim, and see if its (special) or [special]
	trailer := strings.ReplaceAll(cleanedWeb, cleanedDB, "")
	if len(cleanedWeb) != len(trailer) {
		// our dbname is in the webname somewhere
		trailer = strings.TrimSpace(trailer)
		// is is wrapped in round brackets
		if strings.Contains(trailer, "(") && strings.Contains(trailer, ")") {
			return true
		}
		if strings.Contains(trailer, "[") && strings.Contains(trailer, "]") {
			return true
		}
	}
	//fmt.Println(trailer)
	return false
}

func matchAlbums(dbAlbums []*types.Album, webAlbums WebAlbums) []*types.Album {
	//	spew.Dump(dbAlbums)
	//	spew.Dump(webAlbums)
	for i := range webAlbums.Items {
		weGotAMatch := false
		for d := range dbAlbums {
			if stringMatcher(dbAlbums[d].Name, webAlbums.Items[i].Name) {
				dbAlbums[d].Spotify = webAlbums.Items[i].Spotify
				dbAlbums[d].Year = webAlbums.Items[i].Year
				weGotAMatch = true
			}
		}
		if !weGotAMatch && len(dbAlbums) > 0 {
			// new album lets add it as wanted
			newAlbum := types.Album{
				Artist:  dbAlbums[0].Artist,
				Name:    webAlbums.Items[i].Name,
				Year:    webAlbums.Items[i].Year,
				Wanted:  true,
				Format:  "spotify",
				Spotify: webAlbums.Items[i].Spotify,
			}
			dbAlbums = append(dbAlbums, &newAlbum)
		}
	}
	return dbAlbums
}

func lookup(artistID int64, oauth string, artists store.ArtistStore, albums store.AlbumStore, r *http.Request) (matchedAlbums []*types.Album, err error) {
	artist, err := artists.Find(r.Context(), artistID)
	if err != nil {
		return nil, fmt.Errorf("lookup: could not find artist in db: %s", err.Error())
	}
	// perform first lookup for artist.
	if artist.Spotify == "" {
		spotifyID, lookupArtistErr := spotifyLookupArtist(artist.Name, oauth)
		if lookupArtistErr != nil {
			return nil, fmt.Errorf("lookup: could not find artist in db: %s", lookupArtistErr.Error())
		}
		artist.Spotify = spotifyID
		updateErr := artists.Update(r.Context(), artist)
		if updateErr != nil {
			return nil, fmt.Errorf("lookup: unable to update artist: %s", updateErr.Error())
		}
	}
	// lookup the db for the albums
	dbAlbums, err := albums.List(r.Context(), artistID, types.Params{})
	if err != nil {
		return nil, fmt.Errorf("lookup: unable to artists albums: %s", err.Error())
	}
	// lookup spotify for the albums
	webAlbums, err := spotifyLookupArtistAlbums(artist.Spotify, oauth)
	if err != nil {
		return nil, fmt.Errorf("lookup: unable to look up artists spotify albums: %s", err.Error())
	}

	matchedAlbums = matchAlbums(dbAlbums, webAlbums)
	// lets write all this info back to the db
	for _, matchedAlbum := range matchedAlbums {
		if matchedAlbum.ID != 0 {
			// update an album
			_ = albums.Update(r.Context(), matchedAlbum)
		} else {
			// add a new album
			_ = albums.Create(r.Context(), matchedAlbum)
		}
	}
	return matchedAlbums, nil
}

func LookupSingleArtist(artists store.ArtistStore, albums store.AlbumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		artistID, err := strconv.ParseInt(chi.URLParam(r, "artist"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse artist id")
			return
		}
		matchedAlbums, matchErr := lookup(artistID, oauth, artists, albums, r)
		if matchErr != nil {
			render.BadRequest(w, matchErr)
			logger.FromRequest(r).
				WithError(matchErr).
				WithField("artistID", artistID).
				Errorln("cannot match artist")
			return
		}
		render.JSON(w, matchedAlbums, 200)
	}
}

func LookupAllArtists(artistStore store.ArtistStore, albumStore store.AlbumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		id, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		artists, err := artistStore.List(ctx, id, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot retrieve list")
		}
		for _, artist := range artists {
			match, lookupErr := lookup(artist.ID, oauth, artistStore, albumStore, r)
			if lookupErr != nil {
				logger.FromRequest(r).
					WithError(lookupErr).
					WithField("name", artist.Name).
					Errorf("no match for artist")
			} else {
				fmt.Printf("%s:%d,", artist.Name, len(match))
			}
		}
	}
}

// HandleFind returns an http.HandlerFunc that writes the json-encoded artist details to the response body.
func HandleFind(artists store.ArtistStore) http.HandlerFunc {
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

		render.JSON(w, artist, 200)
	}
}

func HandleFindByName(artists store.ArtistStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		project, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}

		name := chi.URLParam(r, "artist")
		artist, err := artists.FindByName(r.Context(), name)
		if err != nil {
			render.NotFound(w, err)
			logger.FromRequest(r).
				WithError(err).
				WithField("string", name).
				Debugln("artist not found")
			return
		}

		if artist.Project != project {
			render.NotFoundf(w, "Not Found")
			logger.FromRequest(r).
				WithField("id", artist.ID).
				WithField("project", project).
				Debugln("project id mismatch")
			return
		}

		render.JSON(w, artist, 200)
	}
}
