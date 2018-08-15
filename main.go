package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"awesomeProject/music"
)
const PORT = "8080"
func main() {
	//err := godotenv.Load() //use .env for env files

	// DB MANAGEMENT
	//db, err := mgo.Dial("localhost")
	//
	//if err != nil {
	//	log.Fatal("cannot dial mongo", err)
	//}
	//defer db.Close()

	//handleInsert(db)

	// SET VARS FOR SPOTIFY
	//config := &clientcredentials.Config{
	//	ClientID:     os.Getenv("SPOTIFY_ID"),
	//	ClientSecret: os.Getenv("SPOTIFY_SECRET"),
	//	TokenURL:     spotify.TokenURL,
	//}

	//client, err := getSpotifyClient(config)
	//
	//
	//results, err := client.GetPlaylistTracks("iplasmic", "2MzgGP8HtsxlZXrIQVSH4g")
	//
	//for _, item := range results.Tracks {
	//	fmt.Println(item.Track.Name)
	//}


	router := music.NewRouter()

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":" + PORT, handlers.CORS(allowedOrigins, allowedMethods)(router)))

}




//func handleInsert(db *mgo.Session){
//	db = db.Copy()
//
//	var s song
//	s.ID = bson.NewObjectId()
//	s.Edited = time.Now()
//
//	var playlist playlist
//
//	playlist.ID = bson.NewObjectId()
//	playlist.Songs = append(playlist.Songs, s.ID)
//
//	err := db.DB("songvc").C("songs").Insert(&s)
//
//	if err != nil{
//		print(err)
//	}
//
//	err = db.DB("songvc").C("playlists").Insert(&playlist)
//
//	if err != nil{
//		print(err)
//	}
//
//}