package main

import (
	"fmt"
	"math"
)

//primeNumber is 素数列挙 int型のchannel
func prinmeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		for i := 3; i < 100; i++ {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 2; j < l+1; j++ {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
		}
		close(result)
	}()
	return result
}

func main() {
	pn := prinmeNumber()
	// pnが動的に増える。fmt実行回数も動的に増える。
	for n := range pn {
		fmt.Println(n)
	}
}
