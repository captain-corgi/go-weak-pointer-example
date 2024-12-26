package main

import (
	"fmt"
	"runtime"
	"weak"
)

func main() {
	// Original pointer
	origin := 5

	// Trigger GC to clean up any unused memory
	runtime.GC()

	// Strong pointerA
	strongA := &origin
	// Use strongA
	fmt.Println(strongA) // Check value of A.
	// 	Example output: 0x1400010c020

	// Trigger GC to clean up any unused memory
	runtime.GC()

	// Weak pointerA
	weakA := weak.Make(strongA) // Example output: 										0x1400010c020

	// Use weakA
	fmt.Println(weakA)         // Example output (Address of weakA): 					{0x1400010c040}
	fmt.Println(weakA.Value()) // Example output (Value of weakA - Address of strongA):	0x1400010c020

	// Trigger GC to clean up any unused memory.
	// 	Since it used, so this GC should clean up weakA
	runtime.GC()

	// Weak pointerA
	//	This one will return nil
	fmt.Println(weakA)         // Example output (Address of weakA):					{0x1400010c040}
	fmt.Println(weakA.Value()) // Example output (Value of weakA - Address of strongA): nil
}
