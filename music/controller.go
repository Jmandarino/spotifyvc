package music

import (
	"net/http"
	"context"
	"github.com/zmb3/spotify"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2/clientcredentials"
	"os"
	"log"
	"fmt"
)

type Controller struct {
	DBconnection DBconnection
}

func getSpotifyClient() (spotify.Client, error){
	err := godotenv.Load() //use .env for env files

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())

	if err != nil {
		log.Fatal("Token auth failed")
	}

	client := spotify.Authenticator{}.NewClient(token)

	return client, err

}



func (c *Controller) Index(w http.ResponseWriter, r *http.Request){
	//playlists := c.DBconnection.InsertSong()
	fmt.Print("Hello world")
	return
}



