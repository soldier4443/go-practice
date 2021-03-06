package main

import ("code.google.com/p/go-tour/pic")

func Pic(dx, dy int) [][]uint8 {
    ret := make([][]uint8, dx)
    for i := 0; i < dx; i++ {
        ret[i] = make([]uint8, dy)
        for j := 0; j < dy; j++ {
            ret[i][j] = uint8((i+j)/2)
        }
    }
    return ret
}

func main() {
    pic.Show(Pic)
}
