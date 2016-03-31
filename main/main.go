package main

import (
	"fmt"
	. "github.com/mpobrien/pewpewpew"
	"math"
)

func color(r Ray, world Traceable) Vector {
	h := world.Trace(r, 0, math.MaxFloat64)
	if h != nil {
		return Vector{h.Normal.X + 1, h.Normal.Y + 1, h.Normal.Z + 1}.Mult(0.5)
	}

	unitDir := r.Direction.Unit()
	t := .5 * (unitDir.Y + 1)
	return Vector{1, 1, 1}.Mult(1.0 - t).Add(Vector{.5, .7, 1.0}.Mult(t))
}

func main() {
	nx, ny := 200, 100
	fmt.Println("P3")
	fmt.Println(nx, ny)
	fmt.Println(255)
	origin := Vector{0, 0, 0}
	lowerLeft := Vector{-2, -1, -1}
	horizontal := Vector{4, 0, 0}
	vertical := Vector{0, 2, 0}

	world := &World{
		Objects: []Traceable{
			Sphere{0.5, Vector{0, 0, -1}},
			Sphere{100, Vector{0, -100.5, -1}},
		},
	}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)
			r := Ray{origin, lowerLeft.Add(horizontal.Mult(u)).Add(vertical.Mult(v))}

			col := color(r, world)
			ir, ig, ib := int(255.99*col.X), int(255.99*col.Y), int(255.99*col.Z)

			fmt.Println(ir, ig, ib)
		}
	}
}
