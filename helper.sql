SELECT
 artist_id

FROM artists 


SELECT
 album_name ,album_format

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
