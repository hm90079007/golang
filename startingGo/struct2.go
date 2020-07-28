package main

import "fmt"

// Point is point
type Point struct {
	X, Y int
}

// Swap is swapping. *があるとポインタ（参照渡し）
func Swap(p *Point) {
	x, y := p.Y, p.X
	p.X = x
	p.Y = y
}

func main() {
	p := &Point{X: 1, Y: 2}
	Swap(p) //&でPoint型のポインタを渡す
	fmt.Printf("p.X = %d, p.Y = %d\n", p.X, p.Y)
	//値渡しなので、別の構造体がコピーされてSwapが実行。したがってp.Xは元のpの値を出力してる
}
