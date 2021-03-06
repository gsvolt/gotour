package main

import (
	"fmt"
	"math"
)

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	//a = v // v is Vertex, not pointer and doesnt implement Abser

	fmt.Println(a.Abs())

	var i I = T{"hello"}
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

}
