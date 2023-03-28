package echo1

import (
	"fmt"
	"os"
)

func echo() {
	//	start := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	//	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	echo()
}
