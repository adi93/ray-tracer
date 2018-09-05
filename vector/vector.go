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
	INVALID_INDEX     = "Invalid index"
)

type Vector interface {
	Get(i int) coordinate
	Length() float64
	Dimension() int

	Add(v *Vector) (Vector, error)
	Subtract(v *Vector) (Vector, error)
	Multiply(v *Vector) (Vector, error)
	Divide(v *Vector) (Vector, error)

	MultiplyByScalar(f float64) Vector
	DivideByScalar(f float64) (Vector, error)

	SquaredLength() float64
	ConvertToUnitVector()

	Equals(v *Vector) bool
}

type coordinate float64

type AbstractVector struct {
	dimension   int
	coordinates []coordinate
}

func New() Vector {
	return AbstractVector{1, []coordinate{0}}
}

func (av AbstractVector) Dimension() int {
	return av.dimension
}
func NewWithValues(rest ...float64) Vector {
	if len(rest) == 0 {
		return New()
	}
	coordinates := make([]coordinate, len(rest))
	for i, x := range rest {
		coordinates[i] = (coordinate)(x)
	}
	return AbstractVector{len(rest), coordinates}
}

func (av AbstractVector) Add(v *Vector) (Vector, error) {
	if reflect.TypeOf(v) != reflect.TypeOf(av) {
		return nil, errors.New(INVALID_TYPE)
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	coordinates := make([]coordinate, maxDimension)
	for i := 0; i < av.dimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) + vCoordinate
	}
	return AbstractVector{av.dimension, coordinates}, nil
}

func max(x, y int) int {
	return (int)(math.Max((float64)(x), (float64)(y)))
}

func (av AbstractVector) Subtract(v *Vector) (Vector, error) {
	if reflect.TypeOf(v) != reflect.TypeOf(av) {
		return nil, errors.New(INVALID_TYPE)
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	coordinates := make([]coordinate, maxDimension)
	for i := 0; i < av.dimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) - vCoordinate
	}
	return AbstractVector{av.dimension, coordinates}, nil
}

func (av AbstractVector) Multiply(v *Vector) (Vector, error) {
	if reflect.TypeOf(v) != reflect.TypeOf(av) {
		return nil, errors.New(INVALID_TYPE)
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	coordinates := make([]coordinate, maxDimension)
	for i := 0; i < av.dimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) * vCoordinate
	}
	return AbstractVector{av.dimension, coordinates}, nil
}

func (av AbstractVector) Divide(v *Vector) (Vector, error) {
	if reflect.TypeOf(v) != reflect.TypeOf(av) {
		return nil, errors.New(INVALID_TYPE)
	}
	maxDimension := max(av.dimension, (*v).Dimension())
	coordinates := make([]coordinate, maxDimension)
	for i := 0; i < av.dimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) / vCoordinate
	}
	return AbstractVector{av.dimension, coordinates}, nil
}

func (av AbstractVector) Get(i int) coordinate {
	if i >= av.dimension {
		return 0
	}
	return av.coordinates[i]
}

func (av AbstractVector) MultiplyByScalar(c float64) Vector {
	return av
}

func (av AbstractVector) DivideByScalar(f float64) (Vector, error) {
	if f == 0 {
		return nil, errors.New(DIVISION_BY_ZERO)
	}
	return av.MultiplyByScalar(1 / f), nil
}

func (av AbstractVector) Length() float64 {
	var length float64
	for _, x := range av.coordinates {
		length += (float64)(x) * (float64)(x)
	}
	return math.Sqrt(length)
}

func (av AbstractVector) SquaredLength() float64 {
	var length float64
	for _, x := range av.coordinates {
		length += (float64)(x) * (float64)(x)
	}
	return length
}

func (av AbstractVector) ConvertToUnitVector() {
	length := av.Length()
	for i, x := range av.coordinates {
		av.coordinates[i] = coordinate((float64)(x) / length)
	}
	return
}

func (av AbstractVector) Equals(v *Vector) bool {
	if reflect.TypeOf(v) != reflect.TypeOf(av) {
		return false
	}
	if av.dimension != (*v).Dimension() {
		return false
	}
	for i, coordinate := range av.coordinates {
		vCoordinate := (*v).Get(i)
		if vCoordinate != coordinate {
			return false
		}
	}
	return true
}

var _ Vector = (*AbstractVector)(nil)

type PositionalVector interface {
	Vector
	X() coordinate
	Y() coordinate
	Z() coordinate
}

type ColorVector interface {
	Vector
	R() coordinate
	G() coordinate
	B() coordinate
}
