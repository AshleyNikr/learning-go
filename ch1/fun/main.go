package main

import "fmt"

func main() {
	s := [2]string{"hello", "world"}
	for _, e := range s {
		fmt.Println(e)
	}
}
