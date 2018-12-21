package objects

import (
	"github.com/adi93/ray-tracer/ray"
	"github.com/adi93/ray-tracer/vector"
)

type Sphere struct {
	Center vector.Pos3Vector
	Radius float64
	Color
}

func (s Sphere) Hit(r ray.Ray, tMin float64, tMax float64, record HitRecord) bool {
	oc := r.Origin.Subtract(&s.Center)
	a := r.Origin.Dot(&r.Origin)
	b := 2.0 * oc.Dot(&r.Direction)
	c := oc.Dot(&oc) - s.Radius
	descriminant := b*b - 4.0*a*c
	return descriminant > 0
}
