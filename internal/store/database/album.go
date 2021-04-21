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

var _ store.AlbumStore = (*AlbumStore)(nil)

// NewAlbumStore returns a new AlbumStore.
func NewAlbumStore(db *sqlx.DB) *AlbumStore {
	return &AlbumStore{db}
}

// AlbumStore implements a AlbumStore backed by a relational
// database.
type AlbumStore struct {
	db *sqlx.DB
}

// Find finds the album by id.
func (s *AlbumStore) Find(ctx context.Context, id int64) (*types.Album, error) {
	dst := new(types.Album)
	err := s.db.Get(dst, albumSelectID, id)
	return dst, err
}

// Find finds the album by string.
func (s *AlbumStore) FindByName(ctx context.Context, artistID int64, str string) (*types.Album, error) {
	dst := new(types.Album)
	err := s.db.Get(dst, albumSelectName, artistID, str)
	return dst, err
}

// List returns a list of albums.
func (s *AlbumStore) List(ctx context.Context, id int64, opts types.Params) ([]*types.Album, error) {
	dst := []*types.Album{}
	err := s.db.Select(&dst, albumSelect, id)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the album details.
func (s *AlbumStore) Create(ctx context.Context, album *types.Album) error {
	query := albumInsert

	if s.db.DriverName() == POSTGRESSTRING {
		query = albumInsertPg
	}

	query, arg, err := s.db.BindNamed(query, album)
	if err != nil {
		return err
	}

	if s.db.DriverName() == POSTGRESSTRING {
		return s.db.QueryRow(query, arg...).Scan(&album.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	album.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the album details.
func (s *AlbumStore) Update(ctx context.Context, album *types.Album) error {
	query, arg, err := s.db.BindNamed(albumUpdate, album)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the album.
func (s *AlbumStore) Delete(ctx context.Context, album *types.Album) error {
	_, err := s.db.Exec(albumDelete, album.ID)
	return err
}

const albumBase = `
SELECT
 album_id
,album_artist_id
,album_name
,album_desc
,album_format
,album_created
,album_updated
FROM albums
`

const albumSelect = albumBase + `
WHERE album_artist_id = $1
ORDER BY album_name ASC
`

const albumSelectID = albumBase + `
WHERE album_id = $1
`

const albumSelectName = albumBase + `
WHERE album_artist_id = $1
AND album_name LIKE $2
`

const albumDelete = `
DELETE FROM albums
WHERE album_id = $1
`

const albumInsert = `
INSERT INTO albums (
 album_artist_id
,album_name
,album_desc
,album_format
,album_created
,album_updated
) values (
 :album_artist_id
,:album_name
,:album_desc
,:album_format
,:album_created
,:album_updated
)
`

const albumInsertPg = albumInsert + `
RETURNING album_id
`

const albumUpdate = `
UPDATE albums
SET
 album_name    = :album_name
,album_desc    = :album_desc
,album_format  = :album_format
,album_updated = :album_updated
WHERE album_id = :album_id
`
