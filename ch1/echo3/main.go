package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
