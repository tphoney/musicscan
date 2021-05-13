// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package artists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tphoney/musicscan/types"
)

const oauth = "BQBX3tRhADAZEGZ9bCvs8tgnjDlhsPIfpv3fM9R1tmJBVPLFZ0NPJt_r0zXsxXh3sqmZrOB2QUmnEzGLMSMRHgxCYBmNs046ndrg3i2TltF_M-UxkTRHz2efPwr44tSo4sSvpTPr6Q2lMrbhMHIHxd1Rzyxtt0k"

func Test_lookupArtist(t *testing.T) {
	artistName := "abba"

	gotSpotifyID, err := lookupArtist(artistName, oauth)
	assert.Equal(t, "0LcJLqbBmaGUft1e9Mm8HV", gotSpotifyID, "got a match for abba")
	assert.Nil(t, err, "no errors")
}

func Test_lookupArtistAlbums(t *testing.T) {
	artistSpotify := "0LcJLqbBmaGUft1e9Mm8HV"

	albums, err := lookupArtistAlbums(artistSpotify, oauth)
	assert.Equal(t, 13, len(albums.Items), "got albums")
	assert.Nil(t, err, "no errors")
}

func Test_matchAlbums(t *testing.T) {
	dbAlbums := []*types.Album{{
		Name:   "A Moon Shaped Pool",
		Format: "flac",
	}, {
		Name:   "Amnesiac",
		Format: "flac",
	}, {
		Name:   "Hail To The Thief",
		Format: "flac",
	}, {
		Name:   "I Might Be Wrong",
		Format: "flac",
	}, {
		Name:   "In Rainbows",
		Format: "flac",
	}, {
		Name:   "Itch",
		Format: "flac",
	}, {
		Name:   "Kid A",
		Format: "flac",
	}, {
		Name:   "OK Computer",
		Format: "flac",
	}, {
		Name:   "OK Computer OKNOTOK 1997 2017",
		Format: "flac",
	}, {
		Name:   "Pablo Honey",
		Format: "flac",
	}, {
		Name:   "Pyramid Song",
		Format: "flac",
	}, {
		Name:   "The Bends",
		Format: "flac",
	}, {
		Name:   "The King of Limbs",
		Format: "flac",
	}}
	var webAlbums WebAlbums
	webAlbums.Items = append(webAlbums.Items,
		struct {
			Name    string
			Year    string
			Spotify string
		}{Name: "OK Computer OKNOTOK 1997 2017",
			Year:    "2017",
			Spotify: "7gMvgOEXNdLYA6Zv0dOcVe"},
		struct {
			Name    string
			Year    string
			Spotify string
		}{Name: "A Moon Shaped Pool",
			Year:    "2016",
			Spotify: "2ix8vWvvSp2Yo7rKMiWpkg"},
		struct {
			Name    string
			Year    string
			Spotify string
		}{Name: "In Rainbows (Disk 2)",
			Year:    "2007",
			Spotify: "6zTAW5oRuOmxJuUHhcQope"},
		struct {
			Name    string
			Year    string
			Spotify string
		}{Name: "TKOL RMX 1234567",
			Year:    "2011",
			Spotify: "566osTxDsfrtdBxPDMGufx"},
	)

	matchedAlbums := matchAlbums(dbAlbums, webAlbums)
	assert.Equal(t, 14, len(matchedAlbums), "Added 1 extra album")
}

func Test_stringMatcher(t *testing.T) {
	type args struct {
		dbName  string
		webName string
	}
	tests := []struct {
		name        string
		args        args
		wantMatched bool
	}{
		{
			name: "simple match",
			args: args{
				dbName:  "HELLO",
				webName: "hello"},
			wantMatched: true,
		},
		{
			name: "special ()",
			args: args{
				dbName:  "In Rainbows",
				webName: "In Rainbows (Disk 2)"},
			wantMatched: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMatched := stringMatcher(tt.args.dbName, tt.args.webName); gotMatched != tt.wantMatched {
				t.Errorf("stringMatcher() = %v, want %v", gotMatched, tt.wantMatched)
			}
		})
	}
}
