package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/jpeg"
	_ "image/png"
)

func main() {
	filePath := os.Args[1]
	fmt.Printf("Attempting to read image from %s\n", filePath)
	imageFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	_, format, err := image.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Read a %s image \n", format)
	return
}
