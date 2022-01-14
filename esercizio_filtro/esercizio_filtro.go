package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var ch string
	var n int

	rigaCmd := os.Args[1:]
	for i := 0; i < len(rigaCmd); i += 2 {

		ch = rigaCmd[i]
		n, _ = strconv.Atoi(rigaCmd[i+1])
		for i2 := 0; i2 < n; i2++ {
			fmt.Print(ch)
		}
	}
	fmt.Println()
}
