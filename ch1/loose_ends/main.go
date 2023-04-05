package main

import (
	"fmt"
	"math/rand"
)

func main() {
	heads := 0
	tails := 0
	iterations := 20
	for i := 1; i < iterations; i++ {
		switch coinFlip() {
		case "heads":
			heads++
		case "tails":
			tails++
		default:
			fmt.Println("landed on edge!")
		}
	}
	fmt.Printf("Heads: %d\nTails: %d\n", heads, tails)
}

func coinFlip() string {
	num := rand.Int()
	if num%2 == 0 {
		return "heads"
	} else {
		return "tails"
	}
}
