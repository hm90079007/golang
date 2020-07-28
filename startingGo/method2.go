package main

import (
	"fmt"
	"math"
)

// Point is point
type Point struct {
	X, Y int
}

// Distance is.pというレシーバーの定義が必要
func (p *Point) Distance(dp *Point) float64 {
	x, y := dp.X-p.X, dp.Y-p.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	p := &Point{X: 0, Y: 0}
	ans := p.Distance(&Point{X: 1, Y: 1})
	fmt.Println(ans)
}
