package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// Distance 普通函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance Point独占方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func main() {
	p := Point{3, 4}
	q := Point{6, 8}
	fmt.Println(p.Distance(q))
	fmt.Println(Distance(p, q))

	fmt.Printf("%T\n", p.Distance)
}
