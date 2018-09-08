package music

import (
	"testing"
)


func TestTrackNewPlaylist(t *testing.T){

	client, err := getSpotifyClient()

	if err != nil{
		t.Errorf("error")
	}

	output := trackNewPlaylist(client, "iplasmic", "2MzgGP8HtsxlZXrIQVSH4g")

	if output != true{
		t.Errorf("trackNewPlaylist failed to return true")
	}
}