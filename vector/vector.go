/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package vector

import (
	"errors"
	"math"
	"reflect"
)

const (
	INVALID_TYPE      = "Invalid type"
	INVALID_DIMENSION = "Invalid dimension"
	DIVISION_BY_ZERO  = "Divison by zero not allowed"
)

type Vector interface {
	Get(i int) float64
	Length() float64
	Dimension() int

	Add(v *Vector) (Vector, error)
	Subtract(v *Vector) (Vector, error)

	MultiplyByScalar(f float64) Vector
	DivideByScalar(f float64) (Vector, error)

	SquaredLength() float64
	ConvertToUnitVector()

	Equals(v *Vector) bool
	Set(i int, value float64)
}

type abstractVector struct {
	dimension   int
	coordinates []float64
}

func New() Vector {
	return abstractVector{1, []float64{0}}
}

func (av abstractVector) Dimension() int {
	return av.dimension
}

func NewFromArray(rest []float64) Vector {
	return abstractVector{len(rest), rest}
}

func NewWithValues(rest ...float64) Vector {
	if len(rest) == 0 {
		return New()
	}
	coordinates := make([]float64, len(rest))
	for i, x := range rest {
		coordinates[i] = (float64)(x)
	}
	return abstractVector{len(rest), coordinates}
}

func (av abstractVector) Add(v *Vector) (Vector, error) {
	if reflect.TypeOf(*v) != reflect.TypeOf(av) {
		return nil, errors.New(INVALID_TYPE)
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	coordinates := make([]float64, maxDimension)
	for i := 0; i < maxDimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) + vCoordinate
	}
	return abstractVector{maxDimension, coordinates}, nil
}

func max(x, y int) int {
	return (int)(math.Max((float64)(x), (float64)(y)))
}

func (av abstractVector) Subtract(v *Vector) (Vector, error) {
	if reflect.TypeOf(*v) != reflect.TypeOf(av) {
		return nil, errors.New(INVALID_TYPE)
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	coordinates := make([]float64, maxDimension)
	for i := 0; i < maxDimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) - vCoordinate
	}
	return abstractVector{maxDimension, coordinates}, nil
}

func (av abstractVector) Get(i int) float64 {
	if i >= av.dimension {
		return 0
	} else if i < 0 {
		panic("coordinate index cannot be negative")
	}
	return av.coordinates[i]
}

func (av abstractVector) MultiplyByScalar(c float64) Vector {
	coordinates := make([]float64, av.dimension)
	for i := 0; i < len(av.coordinates); i++ {
		coordinates[i] = (float64)((float64)(av.Get(i)) * c)
	}
	return abstractVector{av.dimension, coordinates}
}

func (av abstractVector) DivideByScalar(f float64) (Vector, error) {
	if f == 0 {
		return nil, errors.New(DIVISION_BY_ZERO)
	}
	return av.MultiplyByScalar(1 / f), nil
}

func (av abstractVector) Length() float64 {
	var length float64
	for _, x := range av.coordinates {
		length += (float64)(x) * (float64)(x)
	}
	return math.Sqrt(length)
}

func (av abstractVector) SquaredLength() float64 {
	var length float64
	for _, x := range av.coordinates {
		length += (float64)(x) * (float64)(x)
	}
	return length
}

func (av abstractVector) ConvertToUnitVector() {
	length := av.Length()
	if length == 0 {
		return
	}
	for i, x := range av.coordinates {
		av.coordinates[i] = float64((float64)(x) / length)
	}
	return
}

func (av abstractVector) Set(i int, value float64) {
	if i < av.dimension && i >= 0 {
		av.coordinates[i] = value
	}
}

func (av abstractVector) Equals(v *Vector) bool {
	if reflect.TypeOf(*v) != reflect.TypeOf(av) {
		return false
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	for i := 0; i < maxDimension; i++ {
		if (*v).Get(i) != av.Get(i) {
			return false
		}
	}
	return true
}

var _ Vector = (*abstractVector)(nil)

type positionalVec3 struct {
	abstractVector
}

func PositionVec3() positionalVec3 {
	return positionalVec3{abstractVector{3, []float64{0, 0, 0}}}
}

func PositionVec3WithValues(values ...float64) positionalVec3 {
	return positionalVec3{abstractVector{len(values), values}}
}

func (pv3 positionalVec3) X() float64 {
	return pv3.Get(0)
}

func (pv3 positionalVec3) Y() float64 {
	return pv3.Get(0)
}

func (pv3 positionalVec3) Z() float64 {
	return pv3.Get(0)
}

type ColorVector interface {
	Vector
	R() float64
	G() float64
	B() float64
}
