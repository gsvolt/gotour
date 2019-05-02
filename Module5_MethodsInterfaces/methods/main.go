package main

import (
	"fmt"
	"math"
)

func main() {
	Methods()
	Funcs()

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.MyAbs())

	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
