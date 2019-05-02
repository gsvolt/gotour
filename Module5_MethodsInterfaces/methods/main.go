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

	v = Vertex{3, 4}
	ScaleFunction(&v, 10)
	fmt.Println(AbsFunction(v))

	v = Vertex{3, 4}
	v.Scale(2)
	ScaleFunction(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunction(p, 8)

	fmt.Println(v, p)
}
