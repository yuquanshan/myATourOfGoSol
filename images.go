package main

import "golang.org/x/tour/pic"

import (
	"image"
	"image/color"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
	}
	for i := 0; i < dy; i++{
		for j := 0; j<dx; j++ {
			a[i][j] = uint8(i*j)
		}
	}
	return a
}

type Image struct{
	img [][]uint8
}

func (a Image)ColorModel() color.Model{
	return color.RGBAModel
}

func (a Image)Bounds() image.Rectangle{
	return image.Rect(0,0,len(a.img[0]),len(a.img))
}

func (a Image)At(x,y int) color.Color{
	return color.RGBA{a.img[x][y],a.img[x][y],255,255}
}

func main() {
	m := Image{Pic(100,100)}
	pic.ShowImage(m)
}
