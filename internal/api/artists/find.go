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

type SimilarArtistsResponse struct {
	Artists []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
}

type WebAlbums struct {
	Items []struct {
		Name    string
		Year    string
		Spotify string
	}
}

func spotifyLookupArtist(ctx context.Context, artistName, oauth string) (spotifyID string, err error) {
	urlEncodedArtist := url.QueryEscape(artistName)
	artistURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=artist&limit=10", urlEncodedArtist)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, artistURL, nil)
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
	var artistResponse ArtistResponse
	_ = json.Unmarshal(body, &artistResponse)
	for i := range artistResponse.Artists.Items {
		if stringMatcher(artistResponse.Artists.Items[i].Name, artistName) {
			return artistResponse.Artists.Items[i].ID, nil
		}
	}
	return "", fmt.Errorf("no artists match '%s' on spotify", artistName)
}

func spotifyLookupArtistAlbums(ctx context.Context, artistSpotify, oauth string) (albums WebAlbums, err error) {
	albumURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/albums?include_groups=album&market=us&limit=50&", artistSpotify)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, albumURL, nil)
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

func spotifyLookupArtistSimilar(ctx context.Context, artistSpotify, oauth string) (similar SimilarArtistsResponse, err error) {
	similarArtistURL := fmt.Sprintf("https://api.spotify.com/v1/artists/%s/related-artists", artistSpotify)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, httpErr := http.NewRequestWithContext(ctx, http.MethodGet, similarArtistURL, nil)
	if httpErr != nil {
		return similar, fmt.Errorf("spotifyLookupArtistSimilar: get failed from spotify: %s", httpErr.Error())
	}
	bearer := fmt.Sprintf("Bearer %s", oauth)
	req.Header.Add("Authorization", bearer)
	response, err := client.Do(req)
	if err != nil {
		return similar, fmt.Errorf("spotifyLookupArtistSimilar: get failed from spotify: %s", err.Error())
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return similar, fmt.Errorf("spotifyLookupArtistSimilar: unable to parse response from spotify: %s", err.Error())
	}

	var similarArtistsResponse SimilarArtistsResponse
	jsonErr := json.Unmarshal(body, &similarArtistsResponse)
	if jsonErr != nil {
		return similar, fmt.Errorf("spotifyLookupArtistSimilar: unable to unmarshal response from spotify: %s", jsonErr.Error())
	}
	similar.Artists = similarArtistsResponse.Artists
	return similarArtistsResponse, nil
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
	return false
}

func matchSimilarArtists(ctx context.Context, artistStore store.ArtistStore, webArtists SimilarArtistsResponse) (err error) {
	artists, _ := artistStore.List(ctx, 1, types.Params{})
	for i := range webArtists.Artists {
		foundArtist := false
		for j := range artists {
			if stringMatcher(artists[j].Spotify, webArtists.Artists[i].ID) {
				foundArtist = true
				artists[j].Popularity++
				updateErr := artistStore.Update(ctx, artists[j])
				if updateErr != nil {
					return fmt.Errorf("lookup: unable to update artist: %s", updateErr.Error())
				}
			}
		}
		if !foundArtist {
			createErr := artistStore.Create(ctx, &types.Artist{Project: artists[0].Project, Name: webArtists.Artists[i].Name, Desc: "", Spotify: webArtists.Artists[i].ID})
			if createErr != nil {
				return fmt.Errorf("lookup: unable to create artist: %s", createErr.Error())
			}
		}
	}
	return err
}

func matchAlbums(dbAlbums []*types.Album, webAlbums WebAlbums) []*types.Album {
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

func lookupArtistandAlbums(artistID int64, oauth string, artistStore store.ArtistStore, albums store.AlbumStore, r *http.Request) (matchedAlbums []*types.Album, err error) {
	ctx := r.Context()
	artist, err := artistStore.Find(ctx, artistID)
	if err != nil {
		return nil, fmt.Errorf("lookup: could not find artist in db: %s", err.Error())
	}
	// perform first lookup for artist.
	if artist.Spotify == "" {
		spotifyID, lookupArtistErr := spotifyLookupArtist(ctx, artist.Name, oauth)
		if lookupArtistErr != nil {
			return nil, fmt.Errorf("lookup: could not find artist in db: %s", lookupArtistErr.Error())
		}
		artist.Spotify = spotifyID
		updateErr := artistStore.Update(ctx, artist)
		if updateErr != nil {
			return nil, fmt.Errorf("lookup: unable to update artist: %s", updateErr.Error())
		}
	}

	// lookup the db for the albums
	dbAlbums, err := albums.List(ctx, artistID, types.Params{})
	if err != nil {
		return nil, fmt.Errorf("lookup: unable to artists albums: %s", err.Error())
	}
	// lookup spotify for the albums
	webAlbums, err := spotifyLookupArtistAlbums(ctx, artist.Spotify, oauth)
	if err != nil {
		return nil, fmt.Errorf("lookup: unable to look up artists spotify albums: %s", err.Error())
	}
	matchedAlbums = matchAlbums(dbAlbums, webAlbums)
	// lets write all this info back to the db
	for _, matchedAlbum := range matchedAlbums {
		if matchedAlbum.ID != 0 {
			// update an album
			_ = albums.Update(ctx, matchedAlbum)
		} else {
			// add a new album
			_ = albums.Create(ctx, matchedAlbum)
		}
	}
	return matchedAlbums, nil
}

func lookupSimilarArtists(oauth string, artistList []*types.Artist, artistStore store.ArtistStore, r *http.Request) (err error) {
	ctx := r.Context()
	// now lookup similar artists
	for i := range artistList {
		if artistList[i].Spotify != "" {
			similarArtists, lookupSimilarArtistErr := spotifyLookupArtistSimilar(ctx, artistList[i].Spotify, oauth)
			if lookupSimilarArtistErr != nil {
				return fmt.Errorf("lookup: could not find similar artists in spotify: %s", lookupSimilarArtistErr.Error())
			}
			matchSimilarArtistsErr := matchSimilarArtists(ctx, artistStore, similarArtists)
			if matchSimilarArtistsErr != nil {
				return fmt.Errorf("lookup: could not match similar artists: %s", matchSimilarArtistsErr.Error())
			}
		}
	}
	return nil
}

func LookupSingleArtist(artists store.ArtistStore, albums store.AlbumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		oauth := query.Get("spotify_key")

		artistID, err := strconv.ParseInt(chi.URLParam(r, "artist"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse artist id")
			return
		}
		matchedAlbums, matchErr := lookupArtistandAlbums(artistID, oauth, artists, albums, r)
		if matchErr != nil {
			render.BadRequest(w, matchErr)
			logger.FromRequest(r).
				WithError(matchErr).
				WithField("artistID", artistID).
				Errorln("cannot match artist")
			return
		}
		render.JSON(w, matchedAlbums, http.StatusOK)
	}
}

func LookupAllArtists(artistStore store.ArtistStore, albumStore store.AlbumStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		query := r.URL.Query()
		oauth := query.Get("spotify_key")

		projectID, err := strconv.ParseInt(chi.URLParam(r, "project"), 10, 64)
		if err != nil {
			render.BadRequest(w, err)
			logger.FromRequest(r).
				WithError(err).
				Debugln("cannot parse project id")
			return
		}
		cleanArtists, err := artistStore.List(ctx, projectID, types.Params{})
		if err != nil {
			render.InternalError(w, err)
			logger.FromRequest(r).
				WithError(err).
				Errorf("cannot retrieve list")
		}
		// before we start the lookup lets reset artist popularity, and remove recommended artists
		for _, artist := range cleanArtists {
			if artist.Desc == "" {
				deleteErr := artistStore.Delete(ctx, artist)
				if deleteErr != nil {
					logger.FromRequest(r).
						WithError(deleteErr).
						Errorf("cannot delete recommended artist")
				}
				continue
			}
			artist.Popularity = 0
			updateErr := artistStore.Update(r.Context(), artist)
			if updateErr != nil {
				fmt.Printf("unable to reset popularity: %s", updateErr.Error())
			}
		}
		// list again now we dont have recommended artists
		cleanArtists, _ = artistStore.List(ctx, projectID, types.Params{})
		// lets lookup the artists
		for _, artist := range cleanArtists {
			match, lookupErr := lookupArtistandAlbums(artist.ID, oauth, artistStore, albumStore, r)
			if lookupErr != nil {
				logger.FromRequest(r).
					WithError(lookupErr).
					WithField("name", artist.Name).
					Errorf("no match for artist")
			} else {
				fmt.Printf("'%s':%d,", artist.Name, len(match))
			}
		}
		// lookup similar artists
		fmt.Println("artist/album lookup complete, search for similar artists")
		cleanArtists, _ = artistStore.List(ctx, projectID, types.Params{})
		similarErr := lookupSimilarArtists(oauth, cleanArtists, artistStore, r)
		if similarErr != nil {
			logger.FromRequest(r).
				WithError(similarErr).
				Errorf("cannot lookup similar artists")
		}
		fmt.Println("All Lookups complete")
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

		render.JSON(w, artist, http.StatusOK)
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

		render.JSON(w, artist, http.StatusOK)
	}
}
