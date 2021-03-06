package vector

import (
	"errors"
	"math"
	"testing"
)

func TestDimension(t *testing.T) {
	tests := []struct {
		vector            Vector
		expectedDimension int
	}{
		{Vector{}, 0},
		{Vector{1, []float64{1}}, 1},
		{Vector{4, []float64{1, 2.23, -12, 0}}, 4},
	}
	for _, x := range tests {
		actualDimension := x.vector.Dimension()
		expectedDimension := x.expectedDimension
		if actualDimension != expectedDimension {
			t.Fatalf("Dimension check failed. Expected %d, got %d", expectedDimension, actualDimension)
		}
	}
}

func TestGet(t *testing.T) {
	tests := []struct {
		vector              Vector
		expectedCoordinates []float64
	}{
		{Vector{}, []float64{}},
		{Vector{1, []float64{1}}, []float64{1}},
		{Vector{4, []float64{1, 2.23, -12, 0}}, []float64{1, 2.23, -12, 0}},
	}
	for _, x := range tests {
		actualVector := x.vector
		for i, expectedCoord := range x.expectedCoordinates {
			actualCoord := actualVector.Get(i)
			if expectedCoord != actualCoord {
				t.Fatalf("Failed! Expected coordinate %f, found %f", expectedCoord, actualCoord)
			}
		}
	}

	zeroTests := []struct {
		vector Vector
		index  int
	}{
		{Vector{}, 0},
		{Vector{1, []float64{1}}, 1},
		{Vector{1, []float64{1}}, 5},
		{Vector{2, []float64{1, 2}}, 2},
	}

	for _, test := range zeroTests {
		vector := test.vector
		coord := vector.Get(test.index)
		if coord != 0 {
			t.Fatalf("Failed! Expected zero for out of bound index, got %f", coord)
		}
	}
}

func TestGetPanic(t *testing.T) {
	panicTests := []struct {
		vector     Vector
		panicIndex int
	}{
		{Vector{}, -1},
		{Vector{1, []float64{1}}, -5},
		{Vector{2, []float64{1, 2}}, -2},
	}

	for _, test := range panicTests {
		testGetPanic(test.vector, test.panicIndex, t)
	}
}

func testGetPanic(vector Vector, index int, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		}
	}()

	_ = vector.Get(index)
}

func TestNew(t *testing.T) {
	vector := New()
	if vector.Dimension() != 1 {
		t.Fatalf("Dimension check failed! Expected 1, got %d", vector.Dimension())
	}
	if (float64)(vector.Get(0)) != 0 {
		t.Fatalf("Coordinate check failed! Expected 0, got %v", vector.Get(0))
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 3, 3},
		{-1, 123, 123},
		{-45, -10, -10},
		{-10, -45, -10},
		{0, -45, 0},
	}
	for _, test := range tests {
		actualMax := max(test.a, test.b)
		if actualMax != test.expected {
			t.Fatalf("Failed! Expected %d, got %d", test.expected, actualMax)
		}
	}

}

func TestNewWithValues(t *testing.T) {
	tests := []struct {
		vector         Vector
		expectedVector Vector
	}{
		{NewWithValues(1, 2), Vector{2, []float64{1, 2}}},
		{NewWithValues(1), Vector{1, []float64{1}}},
		{NewWithValues(), Vector{1, []float64{0}}},
	}
	for _, x := range tests {
		actualVec := x.vector
		expectedVec := x.expectedVector
		if actualVec.Dimension() != expectedVec.Dimension() {
			t.Fatalf("Dimension check failed. Expected %d, got %d", expectedVec.Dimension(), actualVec.Dimension())
		}
		for i := 0; i < actualVec.Dimension(); i++ {
			actualCoord := actualVec.Get(i)
			expectedCoord := expectedVec.Get(i)
			if expectedCoord != actualCoord {
				t.Fatalf("Failed! Expected %f, found %f", expectedCoord, actualCoord)
			}
		}
	}
}

func TestNewFromArray(t *testing.T) {
	tests := []struct {
		vector         Vector
		expectedVector Vector
	}{
		{NewFromArray([]float64{1, 2}), Vector{2, []float64{1, 2}}},
		{NewFromArray([]float64{1}), Vector{1, []float64{1}}},
		{NewFromArray([]float64{}), Vector{1, []float64{0}}},
	}
	for i, x := range tests {
		actualVec := x.vector
		expectedVec := x.expectedVector
		if actualVec.Dimension() != expectedVec.Dimension() {
			t.Fatalf("Dimension check failed for test %d. Expected %d, got %d", i+1, expectedVec.Dimension(), actualVec.Dimension())
		}
		for d := 0; d < actualVec.Dimension(); d++ {
			actualCoord := actualVec.Get(d)
			expectedCoord := expectedVec.Get(d)
			if expectedCoord != actualCoord {
				t.Fatalf("Failed for test %d! Expected %f, found %f", i+1, expectedCoord, actualCoord)
			}
		}
	}
}

func TestEquals(t *testing.T) {
	tests := []struct {
		vectorA        Vector
		vectorB        Vector
		expectedResult bool
	}{
		{Vector{1, []float64{0}}, Vector{1, []float64{0}}, true},
		{Vector{1, []float64{0}}, NewWithValues(), true},
		{Vector{1, []float64{1, 2}}, Vector{1, []float64{1, 2}}, true},
		{Vector{2, []float64{1, 2}}, Vector{2, []float64{1, 2}}, true},
		{Vector{2, []float64{1, 2}}, Vector{2, []float64{1, 2}}, true},
		{Vector{2, []float64{1, 2}}, Vector{3, []float64{1, 2, 3}}, false},
		{Vector{2, []float64{1, 2}}, Vector{3, []float64{1, 2, 0}}, false},
		{Vector{2, []float64{1, 2}}, Vector{2, []float64{1, -3}}, false},
		{Vector{1, []float64{1}}, Vector{2, []float64{1, -3}}, false},
		{Vector{4, []float64{1, 3.2, -2, -6.67}}, Vector{4, []float64{1, 3.2, -2, -6.66}}, false},
	}
	for i, test := range tests {
		result := test.vectorA.Equals(&test.vectorB)
		if result != test.expectedResult {
			t.Fatalf("Failed for test %d! Expected %v, found %v", i+1, test.expectedResult, result)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		vectorA     Vector
		vectorB     Vector
		expectedSum Vector
	}{
		{New(), New(), New()},
		{NewWithValues(1), NewWithValues(2), NewWithValues(3)},
		{NewWithValues(1, 2), NewWithValues(3, 4), NewWithValues(4, 6)},
	}
	for i, test := range tests {
		actualSum, err := test.vectorA.Add(&test.vectorB)
		if err != nil {
			t.Fatalf("Failed on test %d with error %v", i+1, err)
		}
		if test.expectedSum.Equals(&actualSum) == false {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedSum, actualSum)
		}
	}
}

func TestAddError(t *testing.T) {
	tests := []struct {
		vectorA Vector
		vectorB Vector
		err     error
	}{
		{NewWithValues(1), NewWithValues(1, 2), errors.New(INVALID_DIMENSION)},
		{NewWithValues(), NewWithValues(1, 2), errors.New(INVALID_DIMENSION)},
	}
	for i, test := range tests {
		actualSum, err := test.vectorA.Add(&test.vectorB)
		if err == nil || !actualSum.Equals(&Vector{}) {
			t.Fatalf("Failed on test %d. Expected an error, but got none", i+1)
		}
		if err.Error() != test.err.Error() {
			t.Fatalf("Failed on test %d, Expected error %v, got %v", i+1, test.err, err)
		}
	}
}

func TestSubtract(t *testing.T) {
	tests := []struct {
		vectorA            Vector
		vectorB            Vector
		expectedDifference Vector
	}{
		{New(), New(), New()},
		{NewWithValues(1), NewWithValues(2), NewWithValues(-1)},
		{NewWithValues(3, 4), NewWithValues(1, 0), NewWithValues(2, 4)},
	}
	for i, test := range tests {
		actualDifference, err := test.vectorA.Subtract(&test.vectorB)
		if err != nil {
			t.Fatalf("Failed on test %d with error %v", i+1, err)
		}
		if test.expectedDifference.Equals(&actualDifference) == false {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedDifference, actualDifference)
		}
	}
}

func TestSubtractError(t *testing.T) {
	tests := []struct {
		vectorA Vector
		vectorB Vector
		err     error
	}{
		{NewWithValues(1), NewWithValues(1, 2), errors.New(INVALID_DIMENSION)},
		{NewWithValues(), NewWithValues(1, 2), errors.New(INVALID_DIMENSION)},
		{NewWithValues(1, 2), NewWithValues(), errors.New(INVALID_DIMENSION)},
	}
	for i, test := range tests {
		actualSum, err := test.vectorA.Subtract(&test.vectorB)
		if err == nil || !actualSum.Equals(&Vector{}) {
			t.Fatalf("Failed on test %d. Expected an error, but got none", i+1)
		}
		if err.Error() != test.err.Error() {
			t.Fatalf("Failed on test %d, Expected error %v, got %v", i+1, test.err, err)
		}
	}
}
func TestUnitVector(t *testing.T) {
	tests := []struct {
		vector         Vector
		expectedVector Vector
	}{
		{New(), New()},
		{NewWithValues(1, 2), NewWithValues(1/math.Sqrt(5), 2/math.Sqrt(5))},
		{NewWithValues(0, 0), NewWithValues(0, 0)},
		{NewWithValues(-1, 2), NewWithValues(-1/math.Sqrt(5), 2/math.Sqrt(5))},
		{NewWithValues(1), NewWithValues(1)},
	}
	for i, test := range tests {
		v := test.vector.UnitVector()
		if !v.Equals(&test.expectedVector) {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedVector, v)
		}
	}
}

func TestMultiplyByScalar(t *testing.T) {
	tests := []struct {
		vector         Vector
		scalar         float64
		expectedVector Vector
	}{
		{New(), 12, New()},
		{New(), 0, New()},
		{NewWithValues(1, 2), 2, NewWithValues(2, 4)},
		{NewWithValues(1), 5, NewWithValues(5)},
	}
	for i, test := range tests {
		actualVector := test.vector.MultiplyByScalar(test.scalar)
		if !actualVector.Equals(&test.expectedVector) {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedVector, actualVector)
		}
	}

}

func TestDivideByScalar(t *testing.T) {
	tests := []struct {
		vector         Vector
		scalar         float64
		expectedVector Vector
	}{
		{New(), 12, New()},
		{NewWithValues(1, 2), 2, NewWithValues(0.5, 1)},
		{NewWithValues(1), 5, NewWithValues(0.2)},
	}
	for i, test := range tests {
		actualVector, err := test.vector.DivideByScalar(test.scalar)
		if err != nil {
			t.Fatalf("Found error %v", err)
		}
		if !actualVector.Equals(&test.expectedVector) {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedVector, actualVector)
		}
	}

	divisionByZeroTests := []Vector{
		New(),
		NewWithValues(1, 2),
	}

	for _, vector := range divisionByZeroTests {
		_, err := vector.DivideByScalar(0)
		if err == nil {
			t.Fatalf("Expected division by zero error, got nothing.")
		}
		if err.Error() != DIVISION_BY_ZERO {
			t.Fatalf("Expected %v error ,found %v", DIVISION_BY_ZERO, err.Error())
		}

	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		vector         Vector
		expectedLength float64
	}{
		{New(), 0},
		{NewWithValues(1, 2), math.Sqrt(5)},
		{NewWithValues(1), 1},
		{Vector{2, []float64{0, 0}}, 0},
		{Vector{2, []float64{3, 4}}, 5},
		{Vector{2, []float64{-3, 4}}, 5},
		{Vector{2, []float64{3, -4}}, 5},
		{Vector{2, []float64{-3, -4}}, 5},
	}
	for i, test := range tests {
		actualLength := test.vector.Length()
		if actualLength != test.expectedLength {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedLength, actualLength)
		}
	}
}

func TestSquaredLength(t *testing.T) {
	tests := []struct {
		vector                Vector
		expectedSquaredLength float64
	}{
		{New(), 0},
		{NewWithValues(1, 2), 5},
		{NewWithValues(1), 1},
		{Vector{2, []float64{0, 0}}, 0},
		{Vector{2, []float64{3, 4}}, 25},
		{Vector{2, []float64{-3, 4}}, 25},
		{Vector{2, []float64{3, -4}}, 25},
		{Vector{2, []float64{-3, -4}}, 25},
	}
	for i, test := range tests {
		actualSquaredLength := test.vector.SquaredLength()
		if actualSquaredLength != test.expectedSquaredLength {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedSquaredLength, actualSquaredLength)
		}
	}
}

func TestSet(t *testing.T) {
	tests := []struct {
		vector Vector
		index  int
		value  float64
	}{
		{NewWithValues(1, 2, 3), 0, 4},
		{NewWithValues(1, 2, 3), 2, 10},
		{NewWithValues(1, 2, 3), 2, -10},
		{NewWithValues(1, 2, 3), 2, -10.1234},
		{NewWithValues(1, 2, 3), 2, 10.1234},
	}

	for i, test := range tests {
		test.vector.Set(test.index, test.value)
		foundValue := test.vector.Get(test.index)
		if foundValue != test.value {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.value, foundValue)
		}
	}
}

func TestSetPanic(t *testing.T) {
	panicTests := []struct {
		vector Vector
		index  int
		value  float64
	}{
		{NewWithValues(), -1, 4},
		{NewWithValues(1), -1, 4},
		{NewWithValues(1, 2, 3), -1, 4},
		{NewWithValues(1, 2, 3), -5, 4},
		{NewWithValues(1, 2, 3), 3, 0},
		{NewWithValues(1, 2, 3), 5, 0},
	}

	for _, test := range panicTests {
		testSetPanic(test.vector, test.index, test.value, t)
	}
}

func testSetPanic(vector Vector, index int, value float64, t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("The code did not panic")
		}
	}()

	vector.Set(index, value)
}

func TestDot(t *testing.T) {
	tests := []struct {
		vectorA Vector
		vectorB Vector
		product float64
		err     error
	}{
		{New(), New(), 0, nil},

		// zero vector test
		{NewWithValues(1, 23, -232, 2), NewWithValues(0, 0, 0, 0), 0, nil},

		// squared length test
		{NewWithValues(1, 2, 3), NewWithValues(1, 2, 3), NewWithValues(1, 2, 3).SquaredLength(), nil},

		//normal test
		{NewWithValues(1, 2, 3), NewWithValues(2, 4, 6), 28, nil},

		// invalid dimension test
		{NewWithValues(1, 2, 3), NewWithValues(2, 4, 6, 5), 28, errors.New(INVALID_DIMENSION)},
	}
	for i, test := range tests {
		actualProduct, err := test.vectorA.Dot(&test.vectorB)
		if test.err == nil {
			if err != nil {
				t.Fatalf("Failed on test %d! Expected no error, but found: %v", i+1, err)
			}
			if actualProduct != test.product {
				t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.product, actualProduct)
			}
		} else if err.Error() != test.err.Error() {
			t.Fatalf("Failed on test %d! Expected error: [%v], but found: [%v]", i+1, test.err, err)
		}
	}
}
