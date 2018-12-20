package main

import (
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/adi93/ray-tracer/ray"
	"github.com/adi93/ray-tracer/vector"
)

func simpleGradient(file *os.File) {
	var str strings.Builder
	str.WriteString("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n")

	lowerLeftCorner := vector.NewPos3VectorFromValues(-2, -1, -0.01)
	horizontal := vector.NewPos3VectorFromValues(4, 0, 0)
	vertical := vector.NewPos3VectorFromValues(0, 2, 0)
	origin := vector.NewPos3Vector()

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u, v := float64(i)/float64(nx), float64(j)/float64(ny)
			tempHor, tempVert := horizontal.MultiplyByScalar(u), vertical.MultiplyByScalar(v)
			ray := ray.NewRay(origin, lowerLeftCorner.Add(&tempHor).Add(&tempVert))

			col := color(ray)
			ir := int(255.99 * col.Get(0))
			ig := int(255.99 * col.Get(1))
			ib := int(255.99 * col.Get(2))
			str.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
		}
	}
	_, err := file.WriteString(str.String())
	checkError(err)
}

func color(r ray.Ray) vector.Pos3Vector {
	unitDirection := r.Direction.UnitVector()
	t := 0.5 * (unitDirection.Y() + 1)
	blueComponent := vector.NewPos3VectorFromValues(0.3, 0, 1.0).MultiplyByScalar(t)
	whiteComponent := vector.NewPos3VectorFromValues(1.0, 1.0, 1.0).MultiplyByScalar(1 - t)
	return blueComponent.Add(&whiteComponent)
}
