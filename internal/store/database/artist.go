// Copyright 2019 Brad Rydzewski. All rights reserved.
// Use of this source code is governed by the Polyform License
// that can be found in the LICENSE.md file.

package database

import (
	"context"

	"github.com/tphoney/musicscan/internal/store"
	"github.com/tphoney/musicscan/types"

	"github.com/jmoiron/sqlx"
)

var _ store.ArtistStore = (*ArtistStore)(nil)

// NewArtistStore returns a new ArtistStore.
func NewArtistStore(db *sqlx.DB) *ArtistStore {
	return &ArtistStore{db}
}

// ArtistStore implements a ArtistStore backed by a relational
// database.
type ArtistStore struct {
	db *sqlx.DB
}

// Find finds the artist by id.
func (s *ArtistStore) Find(ctx context.Context, id int64) (*types.Artist, error) {
	dst := new(types.Artist)
	err := s.db.Get(dst, artistSelectID, id)
	return dst, err
}

// Find finds the artist by string.
func (s *ArtistStore) FindByName(ctx context.Context, str string) (*types.Artist, error) {
	dst := new(types.Artist)
	err := s.db.Get(dst, artistSelectName, str)
	return dst, err
}

// List returns a list of artists.
func (s *ArtistStore) List(ctx context.Context, id int64, opts types.Params) ([]*types.Artist, error) {
	dst := []*types.Artist{}
	err := s.db.Select(&dst, artistSelect, id)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the artist details.
func (s *ArtistStore) Create(ctx context.Context, artist *types.Artist) error {
	query := artistInsert

	if s.db.DriverName() == POSTGRESSTRING {
		query = artistInsertPg
	}

	query, arg, err := s.db.BindNamed(query, artist)
	if err != nil {
		return err
	}

	if s.db.DriverName() == POSTGRESSTRING {
		return s.db.QueryRow(query, arg...).Scan(&artist.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	artist.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the artist details.
func (s *ArtistStore) Update(ctx context.Context, artist *types.Artist) error {
	query, arg, err := s.db.BindNamed(artistUpdate, artist)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the artist.
func (s *ArtistStore) Delete(ctx context.Context, artist *types.Artist) error {
	_, err := s.db.Exec(artistDelete, artist.ID)
	return err
}

const artistBase = `
SELECT
 artist_id
,artist_project_id
,artist_name
,artist_desc
,artist_wanted
,artist_popularity
,artist_spotify
,artist_created
,artist_updated
FROM artists
`

const artistSelect = artistBase + `
WHERE artist_project_id = $1
ORDER BY artist_name ASC
`

const artistSelectID = artistBase + `
WHERE artist_id = $1
`

const artistSelectName = artistBase + `
WHERE artist_name LIKE $1 
`

const artistDelete = `
DELETE FROM artists
WHERE artist_id = $1
`

const artistInsert = `
INSERT INTO artists (
 artist_project_id
,artist_name
,artist_desc
,artist_wanted
,artist_popularity
,artist_spotify
,artist_created
,artist_updated
) values (
 :artist_project_id
,:artist_name
,:artist_desc
,:artist_wanted
,:artist_popularity
,:artist_spotify
,:artist_created
,:artist_updated
)
`

const artistInsertPg = artistInsert + `
RETURNING artist_id
`

const artistUpdate = `
UPDATE artists
SET
 artist_name        = :artist_name
,artist_desc        = :artist_desc
,artist_wanted      = :artist_wanted
,artist_popularity  = :artist_popularity
,artist_spotify     = :artist_spotify
,artist_updated     = :artist_updated
WHERE artist_id     = :artist_id
`
