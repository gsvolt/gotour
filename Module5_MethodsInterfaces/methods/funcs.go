package main

import (
	"fmt"
	"math"
)

func AbsFunction(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Funcs() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
