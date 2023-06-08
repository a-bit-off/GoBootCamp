package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func main() {
	width := 300
	height := 300
	border := 15
	// Создание пустого изображения размером width x height
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Заполнение изображения цветом фона (белый)
	backgroundColor := color.RGBA{255, 255, 255, 255}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, backgroundColor)
		}
	}

	// 2
	for x := border; x < 135; x++ {
		for y := border; y < border+45; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for x := 135; x < 180; x++ {
		for y := 60; y < 105; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for x := border + 45; x < 135; x++ {
		for y := 105; y < 150; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for x := border; x < border+45; x++ {
		for y := 150; y < 195; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for x := border + 45; x < 180; x++ {
		for y := 195; y < 240; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	// 1
	for x := 240; x < width-border; x++ {
		for y := 60; y < 240; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}
	for x := 195; x < 240; x++ {
		for y := border; y < border+45; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	// -
	for x := border; x < width-border; x++ {
		for y := 255; y < height-border; y++ {
			img.Set(x, y, color.RGBA{0, 0, 0, 255})
		}
	}

	// Создание файла для сохранения изображения
	file, err := os.Create("logo.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Сохранение изображения в формате PNG
	err = png.Encode(file, img)
	if err != nil {
		log.Fatal(err)
	}
}
