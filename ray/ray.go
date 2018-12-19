package ray

import (
	"github.com/adi93/ray-tracer/vector"
)

type Ray struct {
	Origin    vector.Pos3Vector
	Direction vector.Pos3Vector
}

func New() Ray {
	return Ray{vector.NewPos3Vector(), vector.NewPos3Vector()}
}

func NewRay(a, b vector.Pos3Vector) Ray {
	return Ray{a, b}
}

func (r Ray) PointAtParameter(t float64) vector.Pos3Vector {
	temp := r.Direction.MultiplyByScalar(t)
	return r.Origin.Add(&temp)
}
