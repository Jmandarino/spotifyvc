package music

import (
	"gopkg.in/mgo.v2"
	"log"
	"gopkg.in/mgo.v2/bson"
)



// SERVER the DB server
const SERVER = "localhost"

// DBNAME the name of the DB instance
const DBNAME = "songvc"

// COLLECTION is the name of the collection in DB
const COLLECTIONSONG = "songs"
const COLLECTIONPLIST = "playlists"
const COLLECTIONVERSION = "versions"

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

	err = c.Insert(&p)

	return err == nil

}

func (db DBconnection) InsertVersion(v version) bool{
	session, err := mgo.Dial(SERVER)

	if err != nil {
		log.Fatal("Can't connec to DB")
	}
	defer session.Close()
	c := session.DB(DBNAME).C(COLLECTIONVERSION)

	err = c.Insert(&v)

	return err == nil

}


func (db DBconnection) GetOrCreateSong(s song) song{
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatal("Can't connec to DB")
	}
	defer session.Close()

	var s_ song

	c := session.DB(DBNAME).C(COLLECTIONSONG)
	count, err := c.Find(bson.M{"songId":s.SongId}).Count()//.One(&s_)

	if err != nil {
		log.Fatal("DB error with GetOrCreateSong")
	}

	if count == 0 {
		s.ID = bson.NewObjectId()
		c.Insert(&s)
		return s
	}
	err = c.Find(bson.M{"songId":s.SongId}).One(&s_)

	return s_
}