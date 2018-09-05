package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	TestRange(1, 2, 3, 4)
}

func TestRange(first int, rest ...int) {
	for i, x := range rest {
		fmt.Printf("%dth index: %d\n", i, x)
	}
}

func Cmain() {
	nx, ny := 200, 100
	file, err := os.Create("test.ppm")
	f := 20.5
	fmt.Println(1 / f)
	if err != nil {
		fmt.Println("Could not create temp file")
		os.Exit(1)
	}
	defer file.Close()

	var str strings.Builder
	str.WriteString("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n")

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float64(i) / float64(nx)
			g := float64(j) / float64(ny)
			b := 0.0
			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)
			str.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
		}
	}
	_, err = file.WriteString(str.String())
	if err != nil {
		fmt.Println("Could not write to temp file")
	}
}
