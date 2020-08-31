package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5)

	switch n {
	case 1:
		fmt.Println("大吉")
	case 2, 3:
		fmt.Println("中吉")
	case 4, 5:
		fmt.Println("小吉")
	case 6:
		fmt.Println("凶")
	}

}
