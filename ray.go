package pewpewpew

import "math"

type Scene struct {
}

type Traceable interface {
	Trace(r Ray, tMin, tMax float64) *Hit
}

type Hit struct {
	Normal   Vector
	HitPoint Vector
	T        float64
}

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Mult(t))
}

type World struct {
	Objects []Traceable
}

func (w *World) Trace(r Ray, tMin, tMax float64) *Hit {
	var closestT = tMax
	var closestHit *Hit
	for _, o := range w.Objects {
		hit := o.Trace(r, tMin, closestT)
		if hit != nil {
			closestHit = hit
			closestT = hit.T
		}
	}
	return closestHit
}

type Sphere struct {
	Radius float64
	Center Vector
}

func (s Sphere) Trace(r Ray, tMin, tMax float64) *Hit {
	oc := r.Origin.Sub(s.Center)
	a := Dot(r.Direction, r.Direction)
	b := 2 * Dot(oc, r.Direction)
	c := Dot(oc, oc) - (s.Radius * s.Radius)
	discr := b*b - 4*a*c
	if discr < 0 {
		return nil
	}
	t := (-b - math.Sqrt(discr)) / (2.0 * a)
	if t < tMax && t > tMin {
		p := r.PointAt(t)
		return &Hit{
			T:        t,
			HitPoint: p,
			Normal:   p.Sub(s.Center).Div(s.Radius),
		}
	}
	t = (-b - math.Sqrt(discr)) / (2.0 * a)
	if t < tMax && t > tMin {
		p := r.PointAt(t)
		return &Hit{
			T:        t,
			HitPoint: p,
			Normal:   p.Sub(s.Center).Div(s.Radius),
		}
	}
	return nil
}
