package main

import (
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/adi93/ray-tracer/objects"
	"github.com/adi93/ray-tracer/ray"
	"github.com/adi93/ray-tracer/vector"
)

func simpleSphere(file *os.File) {
	var str strings.Builder
	str.WriteString("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n")

	lowerLeftCorner := vector.NewPos3VectorFromValues(-2, -1, -2)
	horizontal := vector.NewPos3VectorFromValues(4, 0, 0)
	vertical := vector.NewPos3VectorFromValues(0, 2, 0)
	origin := vector.NewPos3Vector()

	world := objects.World{[]objects.HittableObject{objects.Sphere{vector.NewPos3VectorFromValues(1, 0, 0), 3}}}

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u, v := float64(i)/float64(nx), float64(j)/float64(ny)
			tempHor, tempVert := horizontal.MultiplyByScalar(u), vertical.MultiplyByScalar(v)
			ray := ray.NewRay(origin, lowerLeftCorner.Add(&tempHor).Add(&tempVert))

			col := colorWorld(ray, world)
			ir := int(255.99 * col.Get(0))
			ig := int(255.99 * col.Get(1))
			ib := int(255.99 * col.Get(2))
			str.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
		}
	}
	_, err := file.WriteString(str.String())
	checkError(err)
}

func colorWorld(r ray.Ray, world objects.World) vector.Pos3Vector {
	return vector.Pos3Vector{}
}
