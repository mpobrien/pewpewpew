package pewpewpew

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) String() string {
	return fmt.Sprintf("X: %v Y: %v Z: %v", v.X, v.Y, v.Z)
}

func (v Vector) Equals(v2 Vector) bool {
	return v.X == v2.X && v.Y == v2.Y && v.Z == v2.Z
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.X + v2.X, v.Y + v2.Y, v.Z + v2.Z}
}

func (v Vector) Sub(v2 Vector) Vector {
	return Vector{v.X - v2.X, v.Y - v2.Y, v.Z - v2.Z}
}

func (v Vector) Mult(m float64) Vector {
	return Vector{v.X * m, v.Y * m, v.Z * m}
}

func (v Vector) Product(v2 Vector) Vector {
	return Vector{v.X * v2.X, v.Y * v2.Y, v.Z * v2.Z}
}

func (v Vector) Div(m float64) Vector {
	return Vector{v.X / m, v.Y / m, v.Z / m}
}

func (v Vector) Invert() Vector {
	return Vector{v.X * -1, v.Y * -1, v.Z * -1}
}

// Get the pYthagorean length (or magnitude) of the vector.
func (v Vector) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v Vector) SquaredLength() float64 {
	return (v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z)
}

func (v Vector) Unit() Vector {
	l := v.Length()
	return Vector{v.X / l, v.Y / l, v.Z / l}
}

func Dot(v, v2 Vector) float64 {
	return v.X*v2.X + v.Y*v2.Y + v.Z*v2.Z
}

func Cross(v, v2 Vector) Vector {
	return Vector{
		X: v.Y*v2.Z - v.Z*v2.Y,
		Y: -(v.X*v2.Z - v.Z*v2.X),
		Z: v.X*v2.Y - v.Y*v2.X,
	}
}
