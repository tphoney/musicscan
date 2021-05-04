SELECT
 artist_id, artist_wanted

FROM artists 


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

SELECT
 artist_id
,artist_project_id
,artist_name
,artist_desc
,artist_created
,artist_updated
FROM artists

WHERE artist_name LIKE "10cc"
