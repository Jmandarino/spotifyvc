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
	"path/filepath"
)

type Controller struct {
	DBconnection DBconnection
}

func getSpotifyClient() (spotify.Client, error){
	parent, err := os.Getwd()
	path := os.ExpandEnv(filepath.Join(filepath.Dir(parent), ".env"))
	err = godotenv.Load(path)//use .env for env files

	if err != nil{

	}
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

func trackNewPlaylist(client spotify.Client, userName string,  p_id spotify.ID, )bool{

	results, err := client.GetPlaylistTracks(userName, p_id)

	if err != nil {
		log.Fatal("Couldn't get playlist: %v", p_id)
	}

	for _, item := range results.Tracks{
		println(item.Track.Name)
	}

	return true
}



func (c *Controller) Index(w http.ResponseWriter, r *http.Request){
	//playlists := c.DBconnection.InsertSong()
	fmt.Print("Hello world")
	return
}



