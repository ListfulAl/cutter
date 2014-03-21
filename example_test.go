package cutter

import (
	"fmt"
	"image"
	"log"
	"os"
	_ "testing"
)

func ExampleCrop() {
	f, err := os.Open("fixtures/gopher.jpg")
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}

	var c Config
	c.Height = 500
	c.Width = 500

	cImg, err := Crop(img, c)
	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}
	fmt.Println("cImg dimension:", cImg.Bounds())
	// Output: cImg dimension: (0,0)-(500,500)
}
