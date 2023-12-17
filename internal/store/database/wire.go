// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

// WireSet provides a wire set for this package
var WireSet = wire.NewSet(
	ProvideDatabase,
	ProvideUserStore,
	ProvideProjectStore,
	ProvideMemberStore,
	ProvideArtistStore,
	ProvideAlbumStore,
)

// ProvideDatabase provides a database connection.
func ProvideDatabase(config *types.Config) (*sqlx.DB, error) {
	return Connect(
		config.Database.Driver,
		config.Database.Datasource,
	)
}

// ProvideUserStore provides a user store.
func ProvideUserStore(db *sqlx.DB) store.UserStore {
	switch db.DriverName() {
	case POSTGRESSTRING:
		return NewUserStore(db)
	default:
		return NewUserStoreSync(
			NewUserStore(db),
		)
	}
}

// ProvideProjectStore provides a project store.
func ProvideProjectStore(db *sqlx.DB) store.ProjectStore {
	switch db.DriverName() {
	case POSTGRESSTRING:
		return NewProjectStore(db)
	default:
		return NewProjectStoreSync(
			NewProjectStore(db),
		)
	}
}

// ProvideMemberStore provides a member store.
func ProvideMemberStore(db *sqlx.DB) store.MemberStore {
	switch db.DriverName() {
	case POSTGRESSTRING:
		return NewMemberStore(db)
	default:
		return NewMemberStoreSync(
			NewMemberStore(db),
		)
	}
}

// ProvideArtistStore provides a artist store.
func ProvideArtistStore(db *sqlx.DB) store.ArtistStore {
	switch db.DriverName() {
	case POSTGRESSTRING:
		return NewArtistStore(db)
	default:
		return NewArtistStoreSync(
			NewArtistStore(db),
		)
	}
}

// ProvideAlbumStore provides a album store.
func ProvideAlbumStore(db *sqlx.DB) store.AlbumStore {
	switch db.DriverName() {
	case POSTGRESSTRING:
		return NewAlbumStore(db)
	default:
		return NewAlbumStoreSync(
			NewAlbumStore(db),
		)
	}
}
