package main

import (
	"errors"
	"fmt"
	"image"
	"log"
	"os"

	"image/jpeg"
	"image/png"

	"github.com/coding-girls-sofia/go-image-filters/kernel"
)

func loadImage(filePath string) (image.Image, string, error) {
	imageFile, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	imageData, format, err := image.Decode(imageFile)
	if err != nil {
		return nil, "", err
	}

	return imageData, format, nil
}

func writeImage(imageData image.Image, format string) error {
	writer, err := os.Create(fmt.Sprintf("output.%s", format))
	if err != nil {
		log.Fatal(err)
	}

	switch format {
	case "jpeg":
		return jpeg.Encode(writer, imageData, nil)
	case "png":
		return png.Encode(writer, imageData)
	default:
		return errors.New("Unknown format")
	}
}

func fatalError(message string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, message, a...)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		fatalError("Please provide image file name as first argument")
	}
	filePath := os.Args[1]
	fmt.Printf("Attempting to read image from %s\n", filePath)
	imageData, format, err := loadImage(filePath)
	if err != nil {
		fatalError("could not read file: %s", err)
	}
	fmt.Printf("Read a %s image \n", format)
	fmt.Printf("The size of the image is %s\n", imageData.Bounds().Size())

	k := kernel.New([][]float32{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	})
	resultImage, _ := k.Apply(imageData)

	if err := writeImage(resultImage, format); err != nil {
		fatalError("could not write file: %s", err)
	}

	return
}
