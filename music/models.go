package music

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type song struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	SongId  string        `json:"songId" bson:"songId"`
	Title   string        `json:"title" bson:"title"`
	Artists []string      `json:"artists" bson:"artists"`
	Edited  time.Time     `json:"edited" bson:"edited"`
}

type version struct {
	ID         bson.ObjectId   `json:"id" bson:"_id"`
	ChangeType string          `json:"changeType" bson:"changeType"`
	Songs      []bson.ObjectId `json:"songs" bson:"songs"`
	Edited     time.Time       `json:"edited" bson:"edited"`
}

type playlist struct {
	ID         bson.ObjectId   `json:"id" bson:"_id"`
	Name       string          `json:"name" bson:"name"`
	PlaylistId string          `json:"playlistId" bson:"playlistId"`
	UserId     string          `json:"userId" bson:"userId"`
	Songs      []bson.ObjectId `json:"songs" bson:"songs"`
	Versions   []bson.ObjectId `json:"versions" bson:"versions"`
}

type Playlists []playlist
type Songs []song
