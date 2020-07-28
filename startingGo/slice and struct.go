package main

import (
	"fmt"
)

// Point is
type Point struct{ X, Y int }

func main() {
	ps := make([]Point, 5)
	for i, p := range ps {
		fmt.Println(i, p.X, p.Y)
	}
}
