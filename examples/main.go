package main

import (
	"fmt"
	"time"

	"github.com/makasim/rpsmeter"
)

func main() {
	var m rpsmeter.Meter

	go func() {
		ticker := time.NewTicker(time.Millisecond * 10)
		for {
			select {
			case <-ticker.C:
				m.Record()
			}
		}
	}()

	t := time.NewTicker(time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			fmt.Println(m.LastTenSeconds())
		}
	}
}
