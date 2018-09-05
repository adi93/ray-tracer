package vector

import (
	"testing"
)

func TestDimension(t *testing.T) {
	tests := []struct {
		vector            Vector
		expectedDimension int
	}{
		{AbstractVector{}, 0},
		{AbstractVector{1, []coordinate{1}}, 1},
		{AbstractVector{4, []coordinate{1, 2.23, -12, 0}}, 4},
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
		expectedCoordinates []coordinate
	}{
		{AbstractVector{}, []coordinate{}},
		{AbstractVector{1, []coordinate{1}}, []coordinate{1}},
		{AbstractVector{4, []coordinate{1, 2.23, -12, 0}}, []coordinate{1, 2.23, -12, 0}},
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
		{AbstractVector{}, 0},
		{AbstractVector{1, []coordinate{1}}, 1},
		{AbstractVector{2, []coordinate{1, 2}}, 2},
	}

	for _, test := range zeroTests {
		vector := test.vector
		coord := vector.Get(test.panicIndex)
		if coord != 0 {
			t.Fatalf("Failed! Expected nil for out of bound index, got %f", coord)
		}
	}
}

func TestNewWithValues(t *testing.T) {
	tests := []struct {
		vector         Vector
		expectedVector Vector
	}{
		{NewWithValues(1, 2), AbstractVector{2, []coordinate{1, 2}}},
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