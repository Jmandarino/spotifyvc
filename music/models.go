package music

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

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

type Playlists []playlist