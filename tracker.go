package main

import (
	"fmt"
	"time"
)

// Measures the time that program lasted
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)

	fmt.Printf("%s took %s\n", name, elapsed)
}
