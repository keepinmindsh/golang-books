package tips

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

type Vertex3D struct {
	X, Y, Z float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) AbsPointer() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex3D) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Calculate() {
	v := Vertex{3, 4}
	v2 := Vertex3D{3, 4, 5}
	fmt.Println(v.Abs())
	fmt.Println(Abs(v2))
}
