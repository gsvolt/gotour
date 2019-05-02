package main

type MyFloat float64

// MyAbs Receive would not have worked if MyFloat was in a different package
func (f MyFloat) MyAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
