package main

import "fmt"

// Callback is int型のtを引数として、intを返す関数を定義
type Callback func(i int) int

// Sum is int配列型のintsとCallback型（上で宣言）とcallbackを引数でintを返す関数を定義
func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i //range型なので_がインデックス、iがインデックスに対応した値
	}
	return callback(sum)
}

func main() {
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int { return 1 * 2 },
	)
	fmt.Printf("%d\n", n)

}
