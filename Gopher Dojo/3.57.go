package main

import "fmt"

func swap(n, m int) (int, int) {
	return m, n
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	n, m = swap(n, m)
	fmt.Println(n, m)
}
