package main

import (
	"github.com/younisshah/go-watch/heihachi"
	"time"
)

func main() {

	// Make a new ticker for 10 seconds - make this configurable TODO
	ticker := time.NewTicker(time.Duration(10) * time.Second)
	//make a quit/stop channel
	quit := make(chan struct{})
	func() {
		for {
			select {
			case <- ticker.C:
				go_watch.Watch()
			case <- quit:
				ticker.Stop()
				return
			}
		}
	}()
}