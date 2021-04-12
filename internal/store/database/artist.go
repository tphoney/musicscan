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

var _ store.artistStore = (*artistStore)(nil)

// NewartistStore returns a new artistStore.
func NewartistStore(db *sqlx.DB) *artistStore {
	return &artistStore{db}
}

// artistStore implements a artistStore backed by a relational
// database.
type artistStore struct {
	db *sqlx.DB
}

// Find finds the artist by id.
func (s *artistStore) Find(ctx context.Context, id int64) (*types.artist, error) {
	dst := new(types.artist)
	err := s.db.Get(dst, artistSelectID, id)
	return dst, err
}

// List returns a list of artists.
func (s *artistStore) List(ctx context.Context, id int64, opts types.Params) ([]*types.artist, error) {
	dst := []*types.artist{}
	err := s.db.Select(&dst, artistSelect, id)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the artist details.
func (s *artistStore) Create(ctx context.Context, artist *types.artist) error {
	query := artistInsert

	if s.db.DriverName() == "postgres" {
		query = artistInsertPg
	}

	query, arg, err := s.db.BindNamed(query, artist)
	if err != nil {
		return err
	}

	if s.db.DriverName() == "postgres" {
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
func (s *artistStore) Update(ctx context.Context, artist *types.artist) error {
	query, arg, err := s.db.BindNamed(artistUpdate, artist)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the artist.
func (s *artistStore) Delete(ctx context.Context, artist *types.artist) error {
	_, err := s.db.Exec(artistDelete, artist.ID)
	return err
}

const artistBase = `
SELECT
 artist_id
,artist_project_id
,artist_name
,artist_desc
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

const artistDelete = `
DELETE FROM artists
WHERE artist_id = $1
`

const artistInsert = `
INSERT INTO artists (
 artist_project_id
,artist_name
,artist_desc
,artist_created
,artist_updated
) values (
 :artist_project_id
,:artist_name
,:artist_desc
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
 artist_name    = :artist_name
,artist_desc    = :artist_desc
,artist_updated = :artist_updated
WHERE artist_id = :artist_id
`
