package main

import (
    "code.google.com/p/go-tour/pic"
    "math"
    "image"
    "image/color"
)

type Image struct{
	w int
    h int
}

func (i Image) ColorModel() color.Model {
    return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.w, i.h)
}

func (i Image) At(x, y int) color.Color {
    tx := normalize(x, 0, float32(i.w)) * (1 << 8)
    ty := normalize(y, 0, float32(i.h)) * (1 << 8)
    diff := uint8(diff(tx, ty))
    return color.RGBA{diff, diff, diff, math.MaxUint8}
}

func normalize(x int, min float32, max float32) float32 {
    return float32(x) / (max - min) + min
}

func diff(x, y float32) float32 {
    r := x - y
    if r < 0 {
        return -r
    } else {
        return r
    }
}

func main() {
    m := Image{512, 512}
    pic.ShowImage(m)
}
