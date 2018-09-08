package music

import (
	"gopkg.in/mgo.v2"
	"log"
)



// SERVER the DB server
const SERVER = "localhost"

// DBNAME the name of the DB instance
const DBNAME = "songvc"

// COLLECTION is the name of the collection in DB
const COLLECTIONSONG = "songs"
const COLLECTIONPLIST = "playlists"
const COLLECTVERSION = "versions"

type DBconnection struct {}





func (db DBconnection) InsertSong(song song) bool{
	session, err := mgo.Dial(SERVER)

	if err != nil {
		log.Fatal("Can't connec to DB")
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLLECTIONSONG)

	print(c)
	//item := c.Find(bson.M{"SongId":song.SongId})

	return true
}


func (db DBconnection) InsertPlaylist(p playlist) bool{
	session, err := mgo.Dial(SERVER)

	if err != nil {
		log.Fatal("Can't connec to DB")
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTIONPLIST)

	err = c.Insert(p)

	return err == nil

}