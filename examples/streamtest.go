package main

import (
	"log"
	"time"

	"github.com/processone/mpg123"
)

func main() {
	p, err := mpg123.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	p.Play("https://archive.org/download/testmp3testfile/mpthreetest.mp3")
	// TODO Support channel return in player to optionally wait for
	// player to stop It probably need to be in Player itself to have a
	// channel to listen to state changes.
	time.Sleep(12000 * time.Millisecond)
}
