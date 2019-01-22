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
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Controller struct {
	DBconnection DBconnection
}

func getSpotifyClient() (spotify.Client, error){
	parent, err := os.Getwd()
	path := os.ExpandEnv(filepath.Join(parent, ".env"))
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

func trackNewPlaylist(client spotify.Client, userName string,  p_id spotify.ID, ) (p playlist, err error){

	result, err := client.GetPlaylist(userName, p_id)

	if err != nil {
		log.Fatal("Couldn't get playlist: %v", p_id)
		return p, errors.New("spotify error, couldn't get playlist")
	}

	var p_list playlist
	p_list.ID = bson.NewObjectId()
	p_list.UserId = userName
	p_list.PlaylistId = p_id.String()
	p_list.Name = result.Name

	var ver version

	ver.ID = bson.NewObjectId()
	ver.Edited = time.Now()
	ver.ChangeType = "ADD"

	for _, item := range result.Tracks.Tracks{
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
	DBconnection{}.InsertVersion(ver)
	p_list.Versions = append(p_list.Versions, ver.ID)
	DBconnection{}.InsertPlaylist(p_list)

	return p_list, nil
}



func (c *Controller) Index(w http.ResponseWriter, r *http.Request){
	//playlists := c.DBconnection.InsertSong()
	fmt.Print("Hello world")
	return
}

type PlaylistTrack struct {
	User string `json:"user"`
	PId spotify.ID `json:"pid"`
}

func (c *Controller) TrackPlaylist(w http.ResponseWriter, r *http.Request){
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var postData PlaylistTrack
	err = json.Unmarshal(b, &postData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// check if playlist is tracked, else track it
	plist, err := DBconnection{}.GetPlaylistByPlaylistId(postData.User, postData.PId.String())

	if err == nil {
		//playlist is already found, return playlist data
		output, err := json.Marshal(plist)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("content-type", "application/json")
		w.Write(output)
		return
	}

	client, err := getSpotifyClient()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	plist, err = trackNewPlaylist(client, postData.User, postData.PId)
	//start tracking a playlist

	output, err := json.Marshal(postData)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}



