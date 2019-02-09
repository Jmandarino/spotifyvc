package main

import (
	_ "expvar"
	"fmt"
	"github.com/jmandarino/spotifyvc/music"
	"sync"
)

const maxWorkers = 10

type job struct {
	songID string
}

func doWork(id int, j job) {
	fmt.Printf("worker%d: started %s\n", id, j.songID)
	fmt.Printf("worker%d: completed %s!\n", id, j.songID)
}

func main() {

	// channel for jobs
	jobs := make(chan job)

	// start workers
	wg := &sync.WaitGroup{}
	wg.Add(maxWorkers)
	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			defer wg.Done()

			for j := range jobs {
				doWork(i, j)
			}
		}(i)
	}

	// add jobs
	songs := music.DBconnection{}.GetAllSongs()
	for _, song := range songs {
		name := fmt.Sprintf("job-%d", string(song.SongId))
		fmt.Printf("adding: %s\n", name)
		jobs <- job{name}
	}
	close(jobs)

	// wait for workers to complete
	wg.Wait()
}
