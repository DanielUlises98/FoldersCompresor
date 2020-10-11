package main

import (
	"fmt"
	"time"
)

func timeTrack(start time.Time, name string) {
	elpased := time.Since(start)

	fmt.Printf("%s took %s\n", name, elpased)
}
