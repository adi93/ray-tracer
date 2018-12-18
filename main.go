/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

import (
	"github.com/adi93/ray-tracer/ray"
	"github.com/adi93/ray-tracer/vector"
)

var _ = vector.New()

func main() {
	ray := ray.NewRay(vector.NewPos3VectorFromValues(0, 0, 0), vector.NewPos3VectorFromValues(1, 0, 0))
	temp := ray.PointAtParameter(2)
	fmt.Printf("%v\n", temp)
	fmt.Printf("%v\n", reflect.TypeOf(temp))
}

func main2() {
	nx, ny := 200, 100
	log.Printf("%d %d", nx, ny)
	file, err := os.OpenFile("test.ppm", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error %v", err)
		os.Exit(1)
	}
	defer file.Close()
	var str strings.Builder
	str.WriteString("P3\n" + strconv.Itoa(nx) + " " + strconv.Itoa(ny) + "\n255\n")
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := vector.NewWithValues(float64(i)/float64(nx), float64(j)/float64(ny), 0.2)
			ir := int(255.99 * col.Get(0))
			ig := int(255.99 * col.Get(0))
			ib := int(255.99 * col.Get(0))
			str.WriteString(strconv.Itoa(ir) + " " + strconv.Itoa(ig) + " " + strconv.Itoa(ib) + "\n")
		}
	}
	_, err = file.WriteString(str.String())
	if err != nil {
		fmt.Println("Could not write to temp file")
		log.Fatal(err)
	}
}
