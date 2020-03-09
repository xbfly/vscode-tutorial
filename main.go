package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	x := os.Args[1]
	y := strings.ToUpper(x)

	// for i := 0; i < len(x); i++ {
	// 	y += "!"
	// }
	fmt.Println(y)
}
