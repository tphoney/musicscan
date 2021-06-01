// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package types defines common data structures.
package types

import (
	"time"

	"github.com/tphoney/musicscan/types/enum"
	"gopkg.in/guregu/null.v4"
)

type (
	// Artist stores artist details.
	Artist struct {
		ID      int64  `db:"artist_id"         json:"id"`
		Project int64  `db:"artist_project_id" json:"-"`
		Name    string `db:"artist_name"       json:"name"`
		Desc    string `db:"artist_desc"       json:"desc"`
		Wanted  bool   `db:"artist_wanted"     json:"wanted"`
		Spotify string `db:"artist_spotify"    json:"spotify"`
		Created int64  `db:"artist_created"    json:"created"`
		Updated int64  `db:"artist_updated"    json:"updated"`
	}

	// ArtistInput store details used to create or update a artist.
	ArtistInput struct {
		Name   null.String `json:"name"`
		Desc   null.String `json:"desc"`
		Wanted null.Bool   `json:"wanted"`
	}

	// Album stores album details.
	Album struct {
		ID      int64  `db:"album_id"        json:"id"`
		Artist  int64  `db:"album_artist_id" json:"-"`
		Name    string `db:"album_name"      json:"name"`
		Desc    string `db:"album_desc"      json:"desc"`
		Year    string `db:"album_year"      json:"year"`
		Wanted  bool   `db:"album_wanted"    json:"wanted"`
		Spotify string `db:"album_spotify"    json:"spotify"`
		Format  string `db:"album_format"    json:"format"`
		Created int64  `db:"album_created"   json:"created"`
		Updated int64  `db:"album_updated"   json:"updated"`
	}

	// AlbumInput store details used to create or update a album.
	AlbumInput struct {
		Name   null.String `json:"name"`
		Desc   null.String `json:"desc"`
		Format null.String `json:"format"`
		Wanted null.Bool   `json:"wanted"`
		Year   null.String `json:"year"`
	}

	// Member providers member details.
	Member struct {
		Email   string    `db:"user_email"        json:"email"`
		Project int64     `db:"member_project_id" json:"-"`
		User    int64     `db:"member_user_id"    json:"-"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// Membership stores membership details.
	Membership struct {
		Project int64     `db:"member_project_id" json:"-"`
		User    int64     `db:"member_user_id"    json:"-"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// MembershipInput stores membership details.
	MembershipInput struct {
		Project int64     `db:"member_project_id" json:"project"`
		User    string    `db:"member_user_id"    json:"user"`
		Role    enum.Role `db:"member_role"       json:"role"`
	}

	// Params stores query parameters.
	Params struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}

	// Project stores project details.
	Project struct {
		ID      int64  `db:"project_id"      json:"id"`
		Name    string `db:"project_name"    json:"name"`
		Desc    string `db:"project_desc"    json:"desc"`
		Token   string `db:"project_token"   json:"-"`
		Active  bool   `db:"project_active"  json:"active"`
		Created int64  `db:"project_created" json:"created"`
		Updated int64  `db:"project_updated" json:"updated"`
	}

	BadAlbum struct {
		ArtistName null.String `db:"artist_name" json:"artist_name"`
		AlbumName  null.String `db:"album_name" json:"album_name"`
		Format     null.String `db:"album_format" json:"format"`
	}

	// ProjectInput store user project details used to
	// create or update a project.
	ProjectInput struct {
		Name null.String `json:"name"`
		Desc null.String `json:"desc"`
	}

	// Token stores token  details.
	Token struct {
		Value   string    `json:"access_token"`
		Address string    `json:"uri,omitempty"`
		Expires time.Time `json:"expires_at,omitempty"`
	}

	// User stores user account details.
	User struct {
		ID       int64  `db:"user_id"        json:"id"`
		Email    string `db:"user_email"     json:"email"`
		Password string `db:"user_password"  json:"-"`
		Token    string `db:"user_token"     json:"-"`
		Admin    bool   `db:"user_admin"     json:"admin"`
		Blocked  bool   `db:"user_blocked"   json:"-"`
		Created  int64  `db:"user_created"   json:"created"`
		Updated  int64  `db:"user_updated"   json:"updated"`
		Authed   int64  `db:"user_authed"    json:"authed"`
	}

	// UserInput store user account details used to
	// create or update a user.
	UserInput struct {
		Username null.String `json:"email"`
		Password null.String `json:"password"`
		Admin    null.Bool   `json:"admin"`
	}

	// UserToken stores user account and token details.
	UserToken struct {
		User  *User  `json:"user"`
		Token *Token `json:"token"`
	}
)
