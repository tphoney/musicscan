SELECT * FROM artists 
where artist_name LIKE '%Boston%'


SELECT
    artist_id,
    artist_name,
	artist_popularity
FROM
    artists
WHERE
    artist_desc == ''

ORDER BY artist_popularity DESC

SELECT * FROM artists 
where  artist_desc != ''
ORDER BY artist_popularity DESC

SELECT * FROM artists 
where  artist_desc == ''
ORDER BY artist_popularity DESC

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

SELECT
    artists.artist_name, artists.artist_spotify,
    sum(case when  albums.album_format = 'spotify' then 1 else 0 end) as WantedAlbums,
    sum(case when  albums.album_format = 'flac' then 1 else 0 end) as OwnedAlbums
from
    albums
    INNER JOIN artists on artists.artist_id = albums.album_artist_id
WHERE
    artists.artist_wanted == 1
    AND
    album_name NOT LIKE '%live%'
    AND
    album_name NOT LIKE '%anniversary%'
    AND
    album_name NOT LIKE '%deluxe%'
GROUP BY albums.album_artist_id


UPDATE albums
set album_wanted = 1
WHERE album_format != 'flac'

UPDATE artists
set artist_wanted = 0 
WHERE artist_name IN ("10cc", "2Pac", "Aaliyah", "Abba", "Al Stewart", "Alice Cooper", "Alien Ant Farm", "Andy Williams", "Annie Lennox", "Antonio Vivaldi", "Aretha Franklin", "Bad Company", "Barry White", "Bee Gees", "Beethoven", "Bernard Herrmann", "Bill Evans", "Bill Withers", "Billy Joel", "Bing Crosby", "Black Sabbath", "Bob Dylan", "Boney M", "Boy George and Culture Club", "Brandi Carlile", "Brian Kennedy", "Bryan Adams", "Carole King", "Charlie Parker", "Cher", "Chopin", "Chris Rea", "Chuck Berry", "Claude Debussy", "Cliff Richard", "Clint Mansell", "Cyriak", "D-Ream", "David Bowie", "Dean Martin", "Dire Straits", "Disturbed", "Donna Summer", "Doris Day", "Dr. Feelgood", "Dr. John", "Duke Ellington", "Dusty Springfield", "Earth, Wind & Fire", "Electric Light Orchestra", "Elton John", "Elvis Presley", "Ennio Morricone", "Erasure", "Eric Clapton", "Fats Domino", "Foreigner", "Frank Black and The Catholics", "Frank Sinatra", "Franz Schubert", "Garth Brooks", "George Gershwin", "George Michael", "Giuseppe Verdi", "Goldie Lookin Chain", "Groove Armada", "Hall & Oates", "Hans Zimmer", "Hawkwind", "Herbie Hancock", "Huey Lewis & The News", "INXS", "James Brown", "Janis Joplin", "Jefferson Airplane", "Jerry Lee Lewis", "Jethro Tull", "Jimi Hendrix", "Joe Cocker", "Johann Sebastian Bach", "John Denver", "John Lee Hooker", "John Lennon", "John Martyn", "John Williams", "Johnny Cash", "Kylie Minogue", "Ladysmith Black Mambazo", "Linkin Park", "Lionel Richie", "Lou Reed", "Louis Armstrong", "Lynyrd Skynyrd", "Madonna", "Miles Davis", "Ministry", "Mozart", "Nancy Sinatra", "Nat King Cole", "Neil Diamond", "Neil Young", "New York Dolls", "Nickelback", "Nils Lofgren", "Peter Tchaikovsky", "Poco", "Primus", "Prince", "Public Enemy", "Queen", "Queen + Adam Lambert", "Raindance", "Richard Wagner", "Rod Stewart", "Roxette", "Roy Orbison", "Santana", "Seal", "Spice Girls", "Strauss", "Sugababes", "Take That", "Texas", "The Beach Boys", "The Beautiful South", "The Black Eyed Peas", "The Blues Brothers", "The Boomtown Rats", "The Doors", "The Good, The Bad & The Queen", "The Kinks", "The Monkees", "The Muppets", "The Pogues", "The Police", "The Rolling Stones", "The Stranglers", "The Supremes", "The Velvet Underground", "The Who", "The Wombles", "The Zombies", "Tina Turner", "Travis", "Van Halen", "Van Morrison", "Vangelis", "Wham!", "Whitesnake", "William S. Burroughs", "Willie Nelson"
);

SELECT artist_name FROM artists WHERE artist_wanted = 0;

DELETE FROM artists WHERE artist_desc = '';

UPDATE artists set artist_spotify = "";