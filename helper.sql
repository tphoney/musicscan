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
    albums.album_format,
    albums.album_wanted
from
    albums
    INNER JOIN artists on artists.artist_id = albums.album_artist_id
WHERE
    albums.album_format != 'flac' AND albums.album_format != 'spotify'

SELECT
    artists.artist_name,
    albums.album_name,
    albums.album_format,
    albums.album_wanted
from
    albums
    INNER JOIN artists on artists.artist_id = albums.album_artist_id
WHERE
    albums.album_format == 'flac'

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
    albums.album_year LIKE '20%'
    AND
    artists.artist_wanted == 1
    AND
    album_name NOT LIKE '%live%'
    AND
    album_name NOT LIKE '%anniversary%'
    AND
    album_name NOT LIKE '%deluxe%'
ORDER BY album_year DESC

UPDATE albums
set album_wanted = 1
WHERE album_format != 'flac'

UPDATE artists
set artist_wanted = 0 
WHERE artist_name IN ("Al Stewart", "Andy Williams", "Antonio Vivaldi", "Aretha Franklin", "Beethoven", "Bernard Herrmann", "Bill Evans", "Bing Crosby", "Bob Dylan", "Bryan Adams",
"Charlie Parker", "Chopin", "Chris Rea", "Clint Mansell",
"Claude Debussy", "Cliff Richard", "Dean Martin", "Doris Day", "Dr. Feelgood", "Duke Ellington", "Elvis Presley", "Ennio Morricone", "Erasure", "Eric Clapton", "Foreigner", "Frank Sinatra", "Franz Schubert", "George Gershwin", 
"Giuseppe Verdi", "Groove Armada", "Hall & Oates", "Hans Zimmer", "Hawkwind", "James Brown", "Jimi Hendrix",
"Johann Sebastian Bach", "John Williams", "John Martyn", "Johnny Cash", "Kylie Minogue", "Ladysmith Black Mambazo", "Linkin Park", "Lionel Richie", "Lou Reed", "Louis Armstrong", "Lynyrd Skynyrd", "Madonna", "Miles Davis", "Ministry", "Mozart", "Nancy Sinatra", "Nat King Cole", "Neil Diamond", "Neil Young", "New York Dolls", "Nickelback", "Nils Lofgren", "Prince", "Richard Wagner", "Rod Stewart", "Roxette",
"Peter Tchaikovsky","Santana", "Strauss", "The Beach Boys", "The Black Eyed Peas", "Van Morrison", "Vangelis", 
"Willie Nelson");

SELECT artist_name FROM artists WHERE artist_wanted = 0;