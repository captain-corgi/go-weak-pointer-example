package main

import (
	"fmt"
	"runtime"
	"weak"
)

func main() {
	origin := 5

	tmpCount := weak.Make(&origin)
	for {
		if origin == 0 {
			break
		}

		fmt.Println(*(tmpCount.Value()))

		origin--
	}

	runtime.GC()
	fmt.Println(tmpCount.Value())
}
