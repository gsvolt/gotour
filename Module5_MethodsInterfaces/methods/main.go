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
}
