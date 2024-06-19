package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	// for init; con; add {
	// for {}
	for i := 1; i < len(os.Args); i++ {
		s = os.Args[i]
		fmt.Println(i, s)
	}
}

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }
