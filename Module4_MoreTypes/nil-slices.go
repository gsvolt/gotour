package main

import "fmt"

func main() {
    var s []int
    fmt.Printf("Value=%v\tType=%T \n", s, s)
    fmt.Println("s=", s, "length=", len(s), "capacity=", cap(s))
    if s == nil {
        fmt.Println("s is nil!")
    }
    fmt.Println(nil)
}

