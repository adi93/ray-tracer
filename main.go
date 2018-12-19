/*
Copyright 2018 Aditya Harit

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package main

import (
	"log"
	"os"
)

const (
	nx       = 400
	ny       = 200
	testFile = "test.ppm"
)

func main() {
	log.Printf("%d %d", nx, ny)

	file := openFile()
	defer file.Close()

	simpleGradient(file)
}

func openFile() *os.File {
	file, err := os.OpenFile(testFile, os.O_WRONLY, 0600)
	if os.IsNotExist(err) {
		file, err = createFile()
	}
	checkError(err)
	return file
}

func createFile() (*os.File, error) {
	file, err := os.Create(testFile)
	return file, err
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error %v", err)
		os.Exit(1)
	}
}
