package helper

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

func createLog() *log.Logger {
	logFile, err := os.OpenFile("image_cropper.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(logFile, "", log.LstdFlags)

	defer logFile.Close()

	return logger
}

func GetImageBorders(img image.Image) (int, int, int, int) {
	logger := createLog()
	startY := 0
	endY := 0

	startX := 0
	endX := 0

	for i := 0; i < img.Bounds().Size().X; i++ {
		for j := 0; j < img.Bounds().Size().Y; j++ {
			pixel := img.At(i, j)
			r, g, b, _ := pixel.RGBA()
			if r == 0 && g == 0 && b == 0 {
				logger.Println("Black pixel found at:", i, j)

				if startX == 0 {
					startX = j
					startY = i
				}

				if endY == 0 {
					endY = j
				}

				if endX == 0 {
					endX = i
				}

				if endY == j {
					endY++
					log.Default().Println("New black pixel border found at:", i, j)
				}

				if endX == i {
					endX++
					log.Default().Println("New black pixel border found at:", i, j)
				}

			}
		}
	}

	return startX, endX, startY, endY

}

func CropImage(img image.Image, crop image.Rectangle) (image.Image, error) {
	type subImager interface {
		SubImage(r image.Rectangle) image.Image
	}
	simg, ok := img.(subImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}

	return simg.SubImage(crop), nil
}

func WriteImage(img image.Image, name string) error {
	fd, err := os.Create(name)
	if err != nil {
		return err
	}
	defer fd.Close()

	return png.Encode(fd, img)
}

func GetImage(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}
