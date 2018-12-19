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
