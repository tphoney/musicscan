-- name: create-table-albums

CREATE TABLE IF NOT EXISTS albums (
 album_id        INTEGER PRIMARY KEY AUTOINCREMENT
,album_artist_id INTEGER
,album_name      TEXT
,album_desc      TEXT
,album_format    TEXT
,album_created   INTEGER
,album_updated   INTEGER
);

-- name: create-index-album-artist-id

CREATE INDEX IF NOT EXISTS index_album_artist ON albums(album_artist_id);
