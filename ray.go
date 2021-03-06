package pewpewpew

import "math"
import "math/rand"

type Traceable interface {
	Trace(r *Ray, tMin, tMax float64) *Hit
}

type Scatterer interface {
	Scatter(r Ray, h Hit) (Ray, Vector, bool)
}

type Hit struct {
	Normal   Vector
	HitPoint Vector
	T        float64
	Scatterer
}

type Diffuse struct {
	Attenuation float64
}

type Metal struct {
	Attenuation float64
}

func reflect(v, normal Vector) Vector {
	return v.Sub(normal.Mult(2 * Dot(v, normal)))
}

func (m Metal) Scatter(r Ray, h Hit) (Ray, Vector, bool) {
	reflected := reflect(r.Direction.Unit(), h.Normal)
	scattered := Ray{h.HitPoint, reflected}
	return scattered, Vector{0.3, 0.5, 0.9}, Dot(scattered.Direction, h.Normal) > 0
}

func randPointInUnitSphere() Vector {
	p := Vector{0, 0, 0}
	s := float64(100.0)
	for s >= 1.0 {
		p = Vector{rand.Float64(), rand.Float64(), rand.Float64()}.Mult(2.0).Sub(Vector{1, 1, 1})
		s = p.SquaredLength()
	}
	return p
}

func (d Diffuse) Scatter(r Ray, h Hit) (Ray, Vector, bool) {
	target := h.HitPoint.Add(h.Normal).Add(randPointInUnitSphere())
	return Ray{h.HitPoint, target.Sub(h.HitPoint)}, Vector{.5, .5, .5}, true
}

type Camera struct {
	LowerLeft, Horizontal, Vertical, Origin Vector
}

func (c Camera) RayToPixel(u, v float64) Ray {
	return Ray{c.Origin, c.LowerLeft.Add(c.Horizontal.Mult(u)).Add(c.Vertical.Mult(v))}
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

func (w *World) Trace(r *Ray, tMin, tMax float64) *Hit {
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
	Radius   float64
	Center   Vector
	Material Scatterer
}

func fmin(x, y float64) float64 {
	if x < y {
		return x
	}
	return y
}

func fmax(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}

func aabb(vMin, vMax Vector, r *Ray, tMin, tMax float64) bool {
	x01, x02 := (vMin.X-r.Origin.X)/r.Direction.X, (vMax.X - r.Origin.X/r.Direction.X)
	t0, t1 := fmin(x01, x02), fmax(x01, x02)
	tMin = fmax(t0, tMin)
	tMax = fmin(t1, tMax)
	if tMax < tMin {
		return false
	}

	y01, y02 := (vMin.Y-r.Origin.Y)/r.Direction.Y, (vMax.Y - r.Origin.Y/r.Direction.Y)
	t0, t1 = fmin(y01, y02), fmax(y01, y02)
	tMin = fmax(t0, tMin)
	tMax = fmin(t1, tMax)
	if tMax < tMin {
		return false
	}

	z01, z02 := (vMin.Z-r.Origin.Z)/r.Direction.Z, (vMax.Z - r.Origin.Z/r.Direction.Z)
	t0, t1 = fmin(z01, z02), fmax(z01, z02)
	tMin = fmax(t0, tMin)
	tMax = fmin(t1, tMax)
	if tMax < tMin {
		return false
	}
	return true
}

func (s Sphere) Trace(r *Ray, tMin, tMax float64) *Hit {
	//rVec := Vector{s.Radius, s.Radius, s.Radius}
	//if !aabb(s.Center.Sub(rVec), s.Center.Add(rVec), r, tMin, tMax) {
	//return nil
	//}
	//check bounding box
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
			T:         t,
			HitPoint:  p,
			Normal:    p.Sub(s.Center).Div(s.Radius),
			Scatterer: s.Material,
		}
	}
	t = (-b - math.Sqrt(discr)) / (2.0 * a)
	if t < tMax && t > tMin {
		p := r.PointAt(t)
		return &Hit{
			T:         t,
			HitPoint:  p,
			Normal:    p.Sub(s.Center).Div(s.Radius),
			Scatterer: s.Material,
		}
	}
	return nil
}
