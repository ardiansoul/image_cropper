package main

import (
	"bytes"
	"image"
	"image_cropper/helper"
	"log"
)

func main() {
	imageData, err := helper.GetImage("samples/image.png")

	if err != nil {
		log.Panic(err)
	}

	img, _, err := image.Decode(bytes.NewReader(imageData))
	if err != nil {
		log.Panic(err)
	}

	startX, endX, startY, endY := helper.GetImageBorders(img)

	croppedImage, err := helper.CropImage(img, image.Rect(startX, startY, endX, endY))
	if err != nil {
		log.Panic(err)
	}

	err = helper.WriteImage(croppedImage, "samples/image_cropped.png")
	if err != nil {
		log.Panic(err)
	}

}
