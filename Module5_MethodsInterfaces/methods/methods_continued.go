package main

type MyFloat float64

func (f MyFloat) MyAbs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}
