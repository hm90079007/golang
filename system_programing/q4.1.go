package main

import (
	"fmt"
	"time"
)

func main() {
	//	fmt.Println(reflect.TypeOf(time.Now()))
	start := time.Now()
	for {
		var stdin string
		fmt.Scan(&stdin)
		if stdin != "" {
			break
		}
	}
	fmt.Println(time.Since(start))
}
