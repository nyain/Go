package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var nyoba string
	fmt.Scan(&nyoba)
	err := WriteToFile("nyoba.txt", nyoba)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToFile(filename string, data string) error {
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
