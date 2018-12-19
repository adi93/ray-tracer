package main

import (
	"os"
	"strconv"
	"strings"
)

import (
	"github.com/adi93/ray-tracer/vector"
)

func simpleGradient(file *os.File) {
	var str strings.Builder
	str.WriteString("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n")
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := vector.NewWithValues(float64(i)/float64(nx), float64(j)/float64(ny), 0.2)
			ir := int(255.99 * col.Get(0))
			ig := int(255.99 * col.Get(1))
			ib := int(255.99 * col.Get(2))
			str.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
		}
	}
	_, err := file.WriteString(str.String())
	checkError(err)
}

func color(r Ray)
