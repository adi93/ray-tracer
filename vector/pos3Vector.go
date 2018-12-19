package vector

type Pos3Vector struct {
	Vector
}

func NewPos3Vector() Pos3Vector {
	return Pos3Vector{Vector{3, []float64{0, 0, 0}}}
}

func NewPos3VectorFromArray(values [3]float64) Pos3Vector {
	a := make([]float64, 3)
	a[0], a[1], a[2] = values[0], values[1], values[2]
	return Pos3Vector{Vector{3, a}}
}

func NewPos3VectorFromValues(x, y, z float64) Pos3Vector {
	a := make([]float64, 3)
	a[0], a[1], a[2] = x, y, z
	return Pos3Vector{Vector{3, a}}
}

func (pv Pos3Vector) MultiplyByScalar(c float64) Pos3Vector {
	return Pos3Vector{pv.Vector.MultiplyByScalar(c)}
}

func (pv Pos3Vector) DivideByScalar(c float64) (Pos3Vector, error) {
	vector, err := pv.Vector.DivideByScalar(c)
	return Pos3Vector{vector}, err
}

func (pv Pos3Vector) Add(v *Pos3Vector) Pos3Vector {
	sum, _ := pv.Vector.Add(&v.Vector)
	return Pos3Vector{sum}
}

func (pv Pos3Vector) Subtract(v *Pos3Vector) Pos3Vector {
	diff, _ := pv.Vector.Subtract(&v.Vector)
	return Pos3Vector{diff}
}

func (pv Pos3Vector) Equals(v *Pos3Vector) bool {
	return pv.Vector.Equals(&v.Vector)
}

func (pv Pos3Vector) Dot(v *Pos3Vector) float64 {
	product, _ := pv.Vector.Dot(&v.Vector)
	return product
}

func (pv Pos3Vector) UnitVector() Pos3Vector {
	return Pos3Vector{pv.Vector.UnitVector()}
}

func (v1 Pos3Vector) Cross(v2 *Pos3Vector) Pos3Vector {
	return Pos3Vector{
		NewWithValues(
			v1.Get(1)*v2.Get(2)-v1.Get(2)*v2.Get(1),
			v1.Get(2)*v2.Get(0)-v1.Get(0)*v2.Get(2),
			v1.Get(0)*v2.Get(1)-v1.Get(1)*v2.Get(0),
		),
	}
}

func (pv Pos3Vector) X() float64 {
	return pv.Get(0)
}

func (pv Pos3Vector) Y() float64 {
	return pv.Get(1)
}

func (pv Pos3Vector) Z() float64 {
	return pv.Get(2)
}
