package main

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunction(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}
