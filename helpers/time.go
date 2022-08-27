package helpers

import (
	"fmt"
	"time"
)

func Timelog(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
