package main

import (
	"fmt"
	"time"

	greeting "github.com/tenntenn/greeting/v2"
)

func main() {
	now := time.Now()
	fmt.Println(greeting.Do(now))
}
