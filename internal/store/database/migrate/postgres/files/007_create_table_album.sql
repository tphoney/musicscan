-- name: create-table-albums

CREATE TABLE IF NOT EXISTS albums (
 album_id       SERIAL PRIMARY KEY
,album_artist_id   INTEGER
,album_name     VARCHAR(250)
,album_desc     VARCHAR(2000)
,album_created  INTEGER
,album_updated  INTEGER
);

-- name: create-index-album-artist-id

CREATE INDEX IF NOT EXISTS index_album_artist ON albums(album_artist_id);
