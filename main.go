package main

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
	"time"
	"golang.org/x/oauth2/clientcredentials"
	"os"
	"github.com/zmb3/spotify"
	"context"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
	err := godotenv.Load() //use .env for env files

	// DB MANAGEMENT
	db, err := mgo.Dial("localhost")

	if err != nil {
		log.Fatal("cannot dial mongo", err)
	}
	defer db.Close()

	handleInsert(db)

	// SET VARS FOR SPOTIFY
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	client, err := getSpotifyClient(config)


	results, err := client.GetPlaylistTracks("iplasmic", "2MzgGP8HtsxlZXrIQVSH4g")

	for _, item := range results.Tracks {
		fmt.Println(item.Track.Name)
	}

}

func getSpotifyClient(config *clientcredentials.Config) (spotify.Client, error){

	token, err := config.Token(context.Background())

	if err != nil {
		log.Fatal("Token auth failed")
	}

	client := spotify.Authenticator{}.NewClient(token)

	return client, err

}


type song struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	SongId string `json:"songId" bson:"songId"`
	Title string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Edited time.Time `json:"edited" bson:"edited"`
}

type playlist struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	UserId string `json:"userId" bson:"userId"`
	PlaylistId string `json:"playlistId" bson:"playlistId"`
	Songs []bson.ObjectId `json:"songs" bson:"songs"`
}

func handleInsert(db *mgo.Session){
	db = db.Copy()

	var s song
	s.ID = bson.NewObjectId()
	s.Edited = time.Now()

	var playlist playlist

	playlist.ID = bson.NewObjectId()
	playlist.Songs = append(playlist.Songs, s.ID)

	err := db.DB("songvc").C("songs").Insert(&s)

	if err != nil{
		print(err)
	}

	err = db.DB("songvc").C("playlists").Insert(&playlist)

	if err != nil{
		print(err)
	}

}