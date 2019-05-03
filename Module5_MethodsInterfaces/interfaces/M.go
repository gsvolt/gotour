package main

import "fmt"

func (t T) M() {
	fmt.Println(t.S)
}

// func (t *T) M() {
// 	fmt.Println(t.S)
// }

func (f F) M() {
	fmt.Println(f)
}
