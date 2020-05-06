package core

import (
	"time"
)

func Now() *time.Time {
	now := time.Now()
	return &now
}

func GetToday() time.Time {
	year, month, day := time.Now().Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func SetInterval(callback func(), interval time.Duration, async bool) chan bool {
	// Setup the ticket and the channel to signal the ending of the interval
	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	// Put the selection in a go routine
	// so that the for loop is none blocking
	go func() {
		for {
			select {
			case <-ticker.C:
				if async {
					// This won't block
					go callback()
				} else {
					// This will block
					callback()
				}
			case <-clear:
				ticker.Stop()
				return
			}
		}
	}()

	// We return the channel so we can pass in
	// a value to it to clear the interval
	return clear
}

func SetTimeout(callback func(), milliseconds int) {
	timeout := time.Duration(milliseconds) * time.Millisecond
	// This spawns a goroutine and therefore does not block
	time.AfterFunc(timeout, callback)
}
