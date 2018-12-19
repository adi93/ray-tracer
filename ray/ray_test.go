/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package ray

import (
	"github.com/adi93/ray-tracer/vector"
	"testing"
)

func TestNew(t *testing.T) {
	ray := New()
	v := vector.NewPos3Vector()
	if !ray.Origin.Equals(&v) {
		t.Fatalf("Expected origin %v, found %v", v, ray.Origin)
	}

	if !ray.Direction.Equals(&v) {
		t.Fatalf("Expected direction %v, found %v", v, ray.Direction)
	}
}

func TestNewRay(t *testing.T) {
	tests := []struct {
		vectorA vector.Pos3Vector
		vectorB vector.Pos3Vector
	}{
		{vector.NewPos3VectorFromValues(0, 0, 0), vector.NewPos3VectorFromValues(1, 0, 0)},
		{vector.NewPos3Vector(), vector.NewPos3Vector()},
	}

	for i, test := range tests {
		ray := NewRay(test.vectorA, test.vectorB)
		if !ray.Origin.Equals(&test.vectorA) {
			t.Fatalf("Failed for test %d! Expected origin %v, found %v", i+1, test.vectorA, ray.Origin)
		}

		if !ray.Direction.Equals(&test.vectorB) {
			t.Fatalf("Failed for test %d! Expected direction %v, found %v", i+1, test.vectorB, ray.Direction)
		}
	}
}

func TestPointAtParameter(t *testing.T) {
	tests := []struct {
		ray   Ray
		point float64
		value vector.Pos3Vector
	}{
		{NewRay(vector.NewPos3VectorFromValues(0, 0, 0), vector.NewPos3VectorFromValues(1, 0, 0)), 2.0, vector.NewPos3VectorFromValues(2, 0, 0)},
		{NewRay(vector.NewPos3VectorFromValues(0, 0, 0), vector.NewPos3VectorFromValues(0, 0, 0)), 2.0, vector.NewPos3VectorFromValues(0, 0, 0)},
		{NewRay(vector.NewPos3Vector(), vector.NewPos3Vector()), 2.0, vector.NewPos3Vector()},
		{NewRay(vector.NewPos3VectorFromValues(1, 1, 1), vector.NewPos3Vector()), 2.0, vector.NewPos3VectorFromValues(1, 1, 1)},
		{NewRay(vector.NewPos3VectorFromValues(1, 1, 1), vector.NewPos3VectorFromValues(1, 1, 1)), 2.0, vector.NewPos3VectorFromValues(3, 3, 3)},
		{NewRay(vector.NewPos3VectorFromValues(1, 1, 1), vector.NewPos3VectorFromValues(1, 2, 3)), -2.0, vector.NewPos3VectorFromValues(-1, -3, -5)},
	}

	for i, test := range tests {
		actual := test.ray.PointAtParameter(test.point)
		if !actual.Equals(&test.value) {
			t.Fatalf("Failed for test %d! Expected %v, found %v", i+1, test.value, actual)
		}
	}
}
