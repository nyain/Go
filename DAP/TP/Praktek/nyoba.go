package main

import (
	"io"
	"log"
	"os"
)

func main() {
	err := WriteToFile("nyoba.txt", "Ampas")
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}
