package main

import (
	"github.com/adi93/ray-tracer/ray"
	"github.com/adi93/ray-tracer/vector"
)

type HitRecord struct {
	t       float64
	p, norm vector.Pos3Vector
}

type HittableObject interface {
	Hit(r ray.Ray, tMin float64, tMax float64, record HitRecord) bool
}
