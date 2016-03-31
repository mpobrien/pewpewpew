package main

import (
	"fmt"
	. "github.com/mpobrien/pewpewpew"
	"log"
	"math"
	"math/rand"
	"os"
)

var l = log.New(os.Stderr, "", 0)

const MaxDepth = 50

func color(r Ray, world Traceable, depth int) Vector {
	h := world.Trace(r, 0, math.MaxFloat64)
	if h != nil {
		if depth < 50 {
			scattered, attenuation, ok := h.Scatter(r, *h)
			if ok {
				return attenuation.Product(color(scattered, world, depth+1))
			}
			return Vector{0, 0, 0}
		} else {
			return Vector{0, 0, 0}
		}
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
	d := Metal{}

	world := &World{
		Objects: []Traceable{
			Sphere{0.5, Vector{0, 0, -1}, d},
			Sphere{100, Vector{0, -100.5, -1}, d},
		},
	}

	camera := Camera{
		Origin:     Vector{0, 0, 0},
		LowerLeft:  Vector{-2, -1, -1},
		Horizontal: Vector{4, 0, 0},
		Vertical:   Vector{0, 2, 0},
	}

	samplesPerPixel := 100

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			pixelColor := Vector{0, 0, 0}
			for s := 0; s < samplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx)
				v := (float64(j) + rand.Float64()) / float64(ny)
				r := camera.RayToPixel(u, v)
				colSample := color(r, world, 0)
				pixelColor.X += colSample.X
				pixelColor.Y += colSample.Y
				pixelColor.Z += colSample.Z
			}
			pixelColor = pixelColor.Div(float64(samplesPerPixel))
			ir, ig, ib := int(255.99*pixelColor.X), int(255.99*pixelColor.Y), int(255.99*pixelColor.Z)
			fmt.Println(ir, ig, ib)
		}
	}
}
