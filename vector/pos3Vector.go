/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
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
func (pv Pos3Vector) X() float64 {
	return pv.Get(0)
}

func (pv Pos3Vector) Y() float64 {
	return pv.Get(1)
}

func (pv Pos3Vector) Z() float64 {
	return pv.Get(2)
}
