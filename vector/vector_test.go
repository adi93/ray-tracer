/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package vector

import (
	"math"
	"testing"
)

func TestDimension(t *testing.T) {
	tests := []struct {
		vector            Vector
		expectedDimension int
	}{
		{abstractVector{}, 0},
		{abstractVector{1, []float64{1}}, 1},
		{abstractVector{4, []float64{1, 2.23, -12, 0}}, 4},
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
		{abstractVector{}, []float64{}},
		{abstractVector{1, []float64{1}}, []float64{1}},
		{abstractVector{4, []float64{1, 2.23, -12, 0}}, []float64{1, 2.23, -12, 0}},
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
		vector     Vector
		panicIndex int
	}{
		{abstractVector{}, 0},
		{abstractVector{1, []float64{1}}, 1},
		{abstractVector{2, []float64{1, 2}}, 2},
	}

	for _, test := range zeroTests {
		vector := test.vector
		coord := vector.Get(test.panicIndex)
		if coord != 0 {
			t.Fatalf("Failed! Expected zero for out of bound index, got %f", coord)
		}
	}
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

func Testmax(t *testing.T) {
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
		{NewWithValues(1, 2), abstractVector{2, []float64{1, 2}}},
		{NewWithValues(1), abstractVector{1, []float64{1}}},
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

func TestEquals(t *testing.T) {
	tests := []struct {
		vectorA        Vector
		vectorB        Vector
		expectedResult bool
	}{
		{New(), New(), true},
		{abstractVector{1, []float64{0}}, NewWithValues(), true},
		{NewWithValues(1, 2), NewWithValues(1, 2), true},
		{NewWithValues(1, 2), abstractVector{2, []float64{1, 2}}, true},
		{NewWithValues(1, 2), abstractVector{2, []float64{1, 2}}, true},
		{NewWithValues(1, 2), abstractVector{3, []float64{1, 2, 3}}, false},
		{NewWithValues(1, 2), abstractVector{3, []float64{1, 2, 0}}, true},
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
		{NewWithValues(1), NewWithValues(1, 2), NewWithValues(2, 2)},
		{NewWithValues(1.123), NewWithValues(1.123, 2.25), NewWithValues(2.246, 2.25)},
		{NewWithValues(), NewWithValues(1, 2), NewWithValues(1, 2)},
	}
	for _, test := range tests {
		actualSum, err := test.vectorA.Add(&test.vectorB)
		if err != nil {
			t.Fatalf("Failed with error %v", err)
		}
		if test.expectedSum.Equals(&actualSum) == false {
			t.Fatalf("Failed! Expected %v, found %v", test.expectedSum, actualSum)
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
		{NewWithValues(1), NewWithValues(1, 2), NewWithValues(0, -2)},
		{NewWithValues(1, 2), NewWithValues(1), NewWithValues(0, 2)},
		{NewWithValues(1.123), NewWithValues(1.123, 2.25), NewWithValues(0, -2.25)},
		{NewWithValues(), NewWithValues(1, 2), NewWithValues(-1, -2)},
	}
	for i, test := range tests {
		actualDifference, err := test.vectorA.Subtract(&test.vectorB)
		if err != nil {
			t.Fatalf("Failed with error %v", err)
		}
		if test.expectedDifference.Equals(&actualDifference) == false {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedDifference, actualDifference)
		}
	}
}

func TestConvertToUnitVector(t *testing.T) {
	tests := []struct {
		vector         Vector
		expectedVector Vector
	}{
		{New(), New()},
		{NewWithValues(1, 2), NewWithValues(1/math.Sqrt(5), 2/math.Sqrt(5))},
		{NewWithValues(1), NewWithValues(1)},
	}
	for i, test := range tests {
		test.vector.ConvertToUnitVector()
		if test.vector.Equals(&test.expectedVector) == false {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.expectedVector, test.vector)
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
		{NewWithValues(1, 2, 3), 3, 0},
	}
	for i, test := range tests {
		test.vector.Set(test.index, test.value)
		foundValue := test.vector.Get(test.index)
		if foundValue != test.value {
			t.Fatalf("Failed on test %d! Expected %v, found %v", i+1, test.value, foundValue)
		}
	}
}
