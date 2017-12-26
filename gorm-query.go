type Launguage struct {
	ID uint `gorm:"primary_key"`
	Name string
}
	
type Movie struct {
	ID uint `gorm:"primary_key"`
	Title string
	LaunguageID uint
	Launguage Launguage
}

type Artist struct {
	ID uint `gorm:"primary_key"`
	Name string
	Movies []Movie `gorm:"many2many:artist_movies"`
}

//Get list of all artist who acted in “english” movies:
var artists []Artist
if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
	Joins("JOIN movies on movies.id=artist_movies.movie_id").
	Joins("JOIN languages on movies.language_id=languages.id").
	Where("languages.name=?", "english").
	Group("artists.id").Preload("Movies").Find(&artists).Error; err != nil {
	log.Fatal(err)
}
	
for _, ar := range artists {
	fmt.Println(ar.Name)
}

/* output
Kamal Hassan
Madhavan
*/


//Get the list of all artists who acted in movie “Nayagan”
var artists []Artist
if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
	Joins("JOIN movies on artist_movies.movie_id=movies.id").Where("movies.title=?", "Nayagan").
	Group("artists.id").Find(&artists).Error; err != nil {
		log.Fatal(err)
}

for _, ar := range artists {
	fmt.Println(ar.Name)
}

/* output
Kamal Hassan
*/


//Get the list of artists who acted in any of the movies “3 idiots”, “Shamitab” and “310 to Yuma”
var artists []Artist

if err = db.Joins("JOIN artist_movies on artist_movies.artist_id=artists.id").
	Joins("JOIN movies on artist_movies.movie_id=movies.id").
	Where("movies.title in (?)", []string{"3 idiots", "Shamitabh", "310 to Yuma"}).
	Group("artists.id").Find(&artists).Error; err != nil {
		log.Fatal(err)
}

for _, ar := range artists {
	fmt.Println(ar.Name)
}

/* output
Madhavan
Aamir Khan
Christian Bale
Russell Crowe
*/