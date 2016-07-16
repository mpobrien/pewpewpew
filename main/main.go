package main

import (
	"fmt"
	. "github.com/mpobrien/pewpewpew"
	"github.com/pkg/profile"
	"log"
	"math"
	"math/rand"
	"os"
	"sync"
)

var l = log.New(os.Stderr, "", 0)

const MaxDepth = 50

func color(r Ray, world Traceable, depth int) Vector {
	h := world.Trace(&r, 0, math.MaxFloat64)
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
	defer profile.Start().Stop()

	nx, ny := 800, 400
	s := NewScene(nx, ny)

	numSpheres := 100
	world := &World{Objects: make([]Traceable, 0, numSpheres)}
	metal := Metal{}
	diffuse := Diffuse{}

	for i := 0; i < 100; i++ {

		obj := Sphere{float64(rand.Intn(3)) + 1, Vector{float64(rand.Intn(100) - 50), float64(rand.Intn(100) - 50), float64(rand.Intn(100) - 50)}, metal}
		if rand.Intn(2) > 0 {
			obj.Material = diffuse
		}
		world.Objects = append(world.Objects, obj)
	}

	/*
		{
			//Sphere{5, Vector{0, 0, -10}, Diffuse{}},
			Sphere{20, Vector{7, 2, -8}, Metal{}},
			//Sphere{5, Vector{1, 1, 2}, Metal{}},
			Sphere{5, Vector{-.4, 0, -1}, Diffuse{}},
			Sphere{10, Vector{0, -100.5, -1}, Diffuse{}},
		},
	*/
	camera := Camera{
		Origin:     Vector{0, 0, 100},
		LowerLeft:  Vector{-2, -1, -1},
		Horizontal: Vector{4, 0, 0},
		Vertical:   Vector{0, 2, 0},
	}

	samplesPerPixel := 100

	numWorkers := 4
	fmt.Println("starting", numWorkers)
	//runtime.GOMAXPROCS(numWorkers)
	subImageHeight := ny / numWorkers

	wg := sync.WaitGroup{}
	for w := 0; w < numWorkers; w++ {

		wg.Add(1)
		go func(id, minJ, maxJ int) {
			fmt.Println("i'm doing ", minJ, maxJ)
			defer wg.Done()
			for j := minJ; j < maxJ; j++ {
				fmt.Println(id, "doing", j)
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
					ir, ig, ib := 255.99*pixelColor.X, 255.99*pixelColor.Y, 255.99*pixelColor.Z
					s.PutPixel(i, ny-j, Vector{ir, ig, ib})
				}
			}
		}(w, subImageHeight*w, (subImageHeight*w)+subImageHeight)
	}
	wg.Wait()
	s.Save("img.png")
}
