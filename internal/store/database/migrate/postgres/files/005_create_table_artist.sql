-- name: create-table-artists

CREATE TABLE IF NOT EXISTS artists (
 artist_id          SERIAL PRIMARY KEY
,artist_project_id  INTEGER
,artist_name        VARCHAR(250)
,artist_desc        VARCHAR(2000)
,artist_created     INTEGER
,artist_updated     INTEGER
);

-- name: create-index-artist-project-id

CREATE INDEX IF NOT EXISTS index_artist_project ON artists(artist_project_id);
