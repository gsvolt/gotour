package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi) // math module doesn't export pi, but Pi, so you get an error.. pi undefined
}
