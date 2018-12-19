package vector

import (
	"errors"
	"math"
	"strconv"
)

const (
	INVALID_DIMENSION  = "Invalid dimension"
	DIVISION_BY_ZERO   = "Divison by zero not allowed"
	INDEX_OUT_OF_BOUND = "Index out of bound"
)

type Vector struct {
	dimension   int
	coordinates []float64
}

func New() Vector {
	return Vector{1, []float64{0}}
}

func (av Vector) Dimension() int {
	return av.dimension
}

func NewFromArray(values []float64) Vector {
	if len(values) == 0 {
		return Vector{1, []float64{0}}
	}
	return Vector{len(values), values}
}

func NewWithValues(values ...float64) Vector {
	if len(values) == 0 {
		return New()
	}
	coordinates := make([]float64, len(values))
	for i, x := range values {
		coordinates[i] = (float64)(x)
	}
	return Vector{len(values), coordinates}
}

func (av Vector) Add(v *Vector) (Vector, error) {
	if av.dimension != (*v).Dimension() {
		return Vector{}, errors.New(INVALID_DIMENSION)
	}
	dimension := av.dimension
	coordinates := make([]float64, dimension)
	for i := 0; i < dimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) + vCoordinate
	}
	return Vector{dimension, coordinates}, nil
}

func max(x, y int) int {
	a := (int)(math.Max((float64)(x), (float64)(y)))
	return a
}

func (av Vector) Subtract(v *Vector) (Vector, error) {
	if av.dimension != (*v).Dimension() {
		return Vector{}, errors.New(INVALID_DIMENSION)
	}
	dimension := av.dimension
	coordinates := make([]float64, dimension)
	for i := 0; i < dimension; i++ {
		vCoordinate := (*v).Get(i)
		coordinates[i] = av.Get(i) - vCoordinate
	}
	return Vector{dimension, coordinates}, nil
}

func (av Vector) Get(i int) float64 {
	if i >= av.dimension {
		return 0
	} else if i < 0 {
		panic(INDEX_OUT_OF_BOUND + ": " + strconv.Itoa(i))
	}
	return av.coordinates[i]
}

func (av Vector) MultiplyByScalar(c float64) Vector {
	coordinates := make([]float64, av.dimension)
	for i := 0; i < len(av.coordinates); i++ {
		coordinates[i] = (float64)((float64)(av.Get(i)) * c)
	}
	return Vector{av.dimension, coordinates}
}

func (av Vector) DivideByScalar(f float64) (Vector, error) {
	if f == 0 {
		return Vector{}, errors.New(DIVISION_BY_ZERO)
	}
	return av.MultiplyByScalar(1 / f), nil
}

func (av Vector) Length() float64 {
	var length float64
	for _, x := range av.coordinates {
		length += (float64)(x) * (float64)(x)
	}
	return math.Sqrt(length)
}

func (av Vector) SquaredLength() float64 {
	var length float64
	for _, x := range av.coordinates {
		length += (float64)(x) * (float64)(x)
	}
	return length
}

func (av Vector) UnitVector() Vector {
	v := Vector{av.dimension, make([]float64, av.dimension)}
	length := av.Length()
	if length == 0 {
		return v
	}
	for i, x := range av.coordinates {
		v.coordinates[i] = float64((float64)(x) / length)
	}
	return v
}

func (av Vector) Set(i int, value float64) {
	if i < av.dimension && i >= 0 {
		av.coordinates[i] = value
	} else {
		panic(INDEX_OUT_OF_BOUND + ": " + strconv.Itoa(i))
	}
}

func (av Vector) Equals(v *Vector) bool {
	if av.dimension != (*v).Dimension() {
		return false
	}
	dimension := av.dimension
	for i := 0; i < dimension; i++ {
		if (*v).Get(i) != av.Get(i) {
			return false
		}
	}
	return true
}

func (av Vector) Dot(v *Vector) (float64, error) {
	if av.dimension != (*v).Dimension() {
		return 0, errors.New(INVALID_DIMENSION)
	}
	var dotProduct float64 = 0
	for i, x := range av.coordinates {
		dotProduct = dotProduct + x*v.Get(i)
	}
	return dotProduct, nil
}
