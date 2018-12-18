/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package vector

type colorVector struct {
	Vector
}

func NewColorVector() colorVector {
	return colorVector{Vector{3, []float64{0, 0, 0}}}
}

func NewColorVectorFromArray(values [3]int) colorVector {
	a := make([]float64, 3)
	a[0] = (float64)(values[0])
	a[1] = (float64)(values[1])
	a[2] = (float64)(values[2])
	return colorVector{Vector{3, a}}
}

func NewColorVectorFromValues(r, g, b int) colorVector {
	a := make([]float64, 3)
	a[0] = (float64)(r)
	a[1] = (float64)(g)
	a[2] = (float64)(b)
	return colorVector{Vector{3, a}}
}

func (cv colorVector) R() int {
	return castToInt(cv.Get(0))
}

func (cv colorVector) G() int {
	return castToInt(cv.Get(1))
}

func (cv colorVector) B() int {
	return castToInt(cv.Get(2))
}

func castToInt(a float64) int {
	return int(a)
}
