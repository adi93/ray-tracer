package main

import (
	"log"
	"os"
)

const (
	nx       = 1440
	ny       = 900
	testFile = "test.ppm"
)

func main() {
	log.Printf("%d %d", nx, ny)

	file := openFile()
	defer file.Close()

	simpleSphere(file)
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
