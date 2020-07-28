package main

import (
	"fmt"
	//	"math"
)

// Point is point
type Point struct{ X, Y int }

// Render is.pというレシーバーの定義が必要
func (p *Point) Render() {
	fmt.Printf("<%d,%d>\n", p.X, p.Y)
}

func main() {
	p := &Point{X: 5, Y: 10}
	p.Render()
}
