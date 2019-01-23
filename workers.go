package main

import (
	"gopkg.in/mgo.v2/bson"
	"awesomeProject/music"
	"fmt"
)

type Controller struct {
	DBconnection music.DBconnection
}


type WorkRequest struct {
	playListId bson.ObjectId
}

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func main() {
	//Connect to db
	playlists, _ := Controller{}.DBconnection.GetAllPlaylists()

	for _, playlist := range playlists {
		fmt.Println(playlist.PlaylistId)
	}
	// https://gist.github.com/congjf/8035830
	// https://gist.github.com/mvmaasakkers/9088420


	//TODO: get all plist names from db
	//work := WorkRequest{playListId: "123"}
	//
	//WorkQueue <- work
	//fmt.Println("Work request queued " + work.playListId.String())
	//https://gobyexample.com/worker-pools
	// http://motyar.blogspot.com/2015/07/job-queue-example-in-golang.html


}
