package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 1 {
			s := "奇数"
			fmt.Printf("%d-%s\n", i, s)
		} else {
			s := "偶数"
			fmt.Printf("%d-%s\n", i, s)
		}
	}

}
