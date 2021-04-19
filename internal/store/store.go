// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

// Package store defines the data storage interfaces.
package store

import (
	"context"

	"github.com/tphoney/musicscan/types"
)

type (
	// AlbumStore defines album data storage.
	AlbumStore interface {
		// Find finds the album by id.
		Find(ctx context.Context, id int64) (*types.Album, error)

		// Find finds the album by string.
		FindByName(ctx context.Context, id int64, name string) (*types.Album, error)

		// List returns a list of albums by artist id.
		List(ctx context.Context, id int64, params types.Params) ([]*types.Album, error)

		// Create saves the album details.
		Create(ctx context.Context, album *types.Album) error

		// Update updates the album details.
		Update(ctx context.Context, album *types.Album) error

		// Delete deletes the album.
		Delete(ctx context.Context, album *types.Album) error
	}

	// ArtistStore defines artist data storage.
	ArtistStore interface {
		// Find finds the artist by id.
		Find(ctx context.Context, id int64) (*types.Artist, error)

		// Find finds the artist by string.
		FindByName(ctx context.Context, name string) (*types.Artist, error)

		// List returns a list of artists by account id.
		List(ctx context.Context, id int64, params types.Params) ([]*types.Artist, error)

		// Create saves the artist details.
		Create(ctx context.Context, artist *types.Artist) error

		// Update updates the artist details.
		Update(ctx context.Context, artist *types.Artist) error

		// Delete deletes the artist.
		Delete(ctx context.Context, artist *types.Artist) error
	}

	// MemberStore defines member data storage.
	MemberStore interface {
		// Find finds the member by project and user id.
		Find(ctx context.Context, project, user int64) (*types.Member, error)

		// List returns a list of members.
		List(ctx context.Context, project int64, params types.Params) ([]*types.Member, error)

		// Create saves the membership details.
		Create(ctx context.Context, membership *types.Membership) error

		// Update updates the membership details.
		Update(ctx context.Context, membership *types.Membership) error

		// Delete deletes the membership.
		Delete(ctx context.Context, project, user int64) error
	}

	// ProjectStore defines project data storage.
	ProjectStore interface {
		// Find finds the project by id.
		Find(ctx context.Context, id int64) (*types.Project, error)

		// FindToken finds the project by token.
		FindToken(ctx context.Context, token string) (*types.Project, error)

		// List returns a list of projects by user.
		List(ctx context.Context, user int64, params types.Params) ([]*types.Project, error)

		// Create saves the project details.
		Create(ctx context.Context, project *types.Project) error

		// Update updates the project details.
		Update(ctx context.Context, project *types.Project) error

		// Delete deletes the project.
		Delete(ctx context.Context, project *types.Project) error
	}

	// UserStore defines user data storage.
	UserStore interface {
		// Find finds the user by id.
		Find(ctx context.Context, id int64) (*types.User, error)

		// FindEmail finds the user by email.
		FindEmail(ctx context.Context, email string) (*types.User, error)

		// FindKey finds the user by unique key (email or id).
		FindKey(ctx context.Context, key string) (*types.User, error)

		// List returns a list of users.
		List(ctx context.Context, params types.Params) ([]*types.User, error)

		// Create saves the user details.
		Create(ctx context.Context, user *types.User) error

		// Update updates the user details.
		Update(ctx context.Context, user *types.User) error

		// Delete deletes the user.
		Delete(ctx context.Context, user *types.User) error

		// Count returns a count of users.
		Count(ctx context.Context) (int64, error)
	}

	// SystemStore defines insternal system metadata storage.
	SystemStore interface {
		// Config returns the system configuration.
		Config(ctx context.Context) *types.Config
	}
)
