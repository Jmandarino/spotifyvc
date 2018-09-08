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
	"gopkg.in/mgo.v2/bson"
	"time"
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

	var p_list playlist
	p_list.ID = bson.NewObjectId()
	p_list.UserId = userName
	p_list.PlaylistId = p_id.String()

	var ver version

	ver.ID = bson.NewObjectId()
	ver.Edited = time.Now()
	ver.ChangeType = "ADD"

	for _, item := range results.Tracks{
		var s song
		s.SongId = item.Track.ID.String()
		for _, artist := range item.Track.Artists{
			s.Artists = append(s.Artists, artist.Name)
		}
		s.Title = item.Track.Name
		song := DBconnection{}.GetOrCreateSong(s)

		p_list.Songs = append(p_list.Songs, song.ID)
		ver.Songs = append(ver.Songs, song.ID)
	}
	p_list.Versions = append(p_list.Versions, ver.ID)
	DBconnection{}.InsertVersion(ver)
	DBconnection{}.InsertPlaylist(p_list)

	return true
}



func (c *Controller) Index(w http.ResponseWriter, r *http.Request){
	//playlists := c.DBconnection.InsertSong()
	fmt.Print("Hello world")
	return
}



