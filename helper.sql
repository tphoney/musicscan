SELECT
 artist_id, artist_name, artist_wanted, artist_spotify

FROM artists 

where artist_name LIKE '%radiohead%'


SELECT
 album_id
,album_artist_id
,album_name
,album_desc
,album_format
,album_year
,album_wanted
,album_created
,album_updated
FROM albums
WHERE 
albums.album_artist_id == 5

SELECT
    artists.artist_name,
    albums.album_name,
    albums.album_format
from
    albums
    INNER JOIN artists on artists.artist_id = albums.album_artist_id
WHERE
    albums.album_format != 'flac'

SELECT
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
    albums.album_year == 2021
    AND
    artists.artist_wanted == 1
