-- name: create-table-artists

CREATE TABLE IF NOT EXISTS artists (
 artist_id          INTEGER PRIMARY KEY AUTOINCREMENT
,artist_project_id  INTEGER
,artist_name        TEXT
,artist_desc        TEXT
,artist_wanted      BOOLEAN DEFAULT 1
,artist_created     INTEGER
,artist_updated     INTEGER
);

-- name: create-index-artist-project-id

CREATE INDEX IF NOT EXISTS index_artist_project ON artists(artist_project_id);
