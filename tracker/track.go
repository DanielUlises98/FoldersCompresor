package tracker

import (
	"fmt"
	"time"
)

//TimeTrack ...  Measures the time that program lasted
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)

	fmt.Printf("%s took %s\n", name, elapsed)
}
