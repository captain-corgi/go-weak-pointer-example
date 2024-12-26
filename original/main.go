package main

import (
	"runtime"
	"weak"
)

type T struct {
	a int
	b int
}

func main() {
	a := new(string)
	println("original:", a)

	// make a weak pointer
	weakA := weak.Make(a)

	runtime.GC()

	// use weakA
	strongA := weakA.Value()
	println("strong:", strongA, a)

	runtime.GC()

	// use weakA again
	strongA = weakA.Value()
	println("strong:", strongA)
}

// Output:
// original: 0x1400010c670
// strong: 0x1400010c670 0x1400010c670
// strong: 0x0
