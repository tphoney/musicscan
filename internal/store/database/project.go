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

var _ store.ProjectStore = (*ProjectStore)(nil)

// NewProjectStore returns a new ProjecttStore.
func NewProjectStore(db *sqlx.DB) *ProjectStore {
	return &ProjectStore{db}
}

// ProjectStore implements a ProjectStore backed by a
// relational database.
type ProjectStore struct {
	db *sqlx.DB
}

// Find finds the project by id.
func (s *ProjectStore) Find(ctx context.Context, id int64) (*types.Project, error) {
	dst := new(types.Project)
	err := s.db.Get(dst, projectSelectID, id)
	return dst, err
}

// Find finds the project by id.
func (s *ProjectStore) FindBadAlbums(ctx context.Context, id int64) ([]*types.BadAlbum, error) {
	dst := []*types.BadAlbum{}
	err := s.db.Select(&dst, projectBadAlbums)
	return dst, err
}

// Find finds the project by id.
func (s *ProjectStore) FindRecommendedArtists(ctx context.Context, id int64) ([]*types.Artist, error) {
	dst := []*types.Artist{}
	err := s.db.Select(&dst, projectRecommendedArtists, id)
	return dst, err
}

// Find finds the project by id.
func (s *ProjectStore) FindWantedAlbums(ctx context.Context, id, year int64) ([]*types.BadAlbum, error) {
	dst := []*types.BadAlbum{}
	err := s.db.Select(&dst, projectWantedAlbums, year)
	return dst, err
}

// FindToken finds the project by token.
func (s *ProjectStore) FindToken(ctx context.Context, token string) (*types.Project, error) {
	dst := new(types.Project)
	err := s.db.Get(dst, projectSelectToken, token)
	return dst, err
}

// List returns a list of projects by user.
func (s *ProjectStore) List(ctx context.Context, user int64, opts types.Params) ([]*types.Project, error) {
	dst := []*types.Project{}
	err := s.db.Select(&dst, projectSelect, user)
	// TODO(bradrydzewski) add limit and offset
	return dst, err
}

// Create saves the project details.
func (s *ProjectStore) Create(ctx context.Context, project *types.Project) error {
	query := projectInsert

	if s.db.DriverName() == POSTGRESSTRING {
		query = projectInsertPg
	}

	query, arg, err := s.db.BindNamed(query, project)
	if err != nil {
		return err
	}

	if s.db.DriverName() == POSTGRESSTRING {
		return s.db.QueryRow(query, arg...).Scan(&project.ID)
	}

	res, err := s.db.Exec(query, arg...)
	if err != nil {
		return err
	}
	project.ID, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

// Update updates the project details.
func (s *ProjectStore) Update(ctx context.Context, project *types.Project) error {
	query, arg, err := s.db.BindNamed(projectUpdate, project)
	if err != nil {
		return err
	}
	_, err = s.db.Exec(query, arg...)
	return err
}

// Delete deletes the project.
func (s *ProjectStore) Delete(ctx context.Context, project *types.Project) error {
	_, err := s.db.Exec(projectDelete, project.ID)
	return err
}

const projectBase = `
SELECT
 project_id
,project_name
,project_desc
,project_token
,project_active
,project_created
,project_updated
FROM projects
`

const projectSelect = projectBase + `
WHERE project_id IN (
  SELECT member_project_id
  FROM members
  WHERE member_user_id = $1
)
ORDER BY project_name
`

const projectSelectID = projectBase + `
WHERE project_id = $1
`

const projectBadAlbums = `
SELECT
    artists.artist_id,
    artists.artist_name,
    albums.album_name,
    albums.album_format,
	albums.album_year
from
    albums
    INNER JOIN artists on artists.artist_id = albums.album_artist_id
WHERE
    albums.album_format != 'flac' AND albums.album_format != 'spotify'
`

const projectWantedAlbums = `
SELECT
	artists.artist_id,
    artists.artist_name,
    albums.album_name,
    albums.album_format,
    albums.album_year
from
    albums
    INNER JOIN artists on artists.artist_id = albums.album_artist_id
WHERE
    albums.album_format == 'spotify'
    AND
    albums.album_year == $1
    AND
    artists.artist_wanted == 1
    AND
    album_name NOT LIKE '%live%'
    AND
    album_name NOT LIKE '%anniversary%'
    AND
    album_name NOT LIKE '%deluxe%'
ORDER BY album_year DESC
`
const projectRecommendedArtists = `
SELECT
    artist_id,
    artist_name,
	artist_spotify,
	artist_popularity
FROM
    artists
WHERE
    artist_desc == ''
AND
	artist_project_id = $1
ORDER BY artist_popularity DESC
`

const projectSelectToken = projectBase + `
WHERE project_token = $1
`

const projectDelete = `
DELETE FROM projects
WHERE project_id = $1
`

const projectInsert = `
INSERT INTO projects (
 project_name
,project_desc
,project_token
,project_active
,project_created
,project_updated
) values (
 :project_name
,:project_desc
,:project_token
,:project_active
,:project_created
,:project_updated
)
`

const projectInsertPg = projectInsert + `
RETURNING project_id
`

const projectUpdate = `
UPDATE projects
SET
 project_name      = :project_name
,project_desc      = :project_desc
,project_active    = :project_active
,project_updated   = :project_updated
WHERE project_id = :project_id
`
