package main

import "golang.org/x/tour/pic"

// Pic creates a picture
// interresting values are:
//	// (x+y)/2
//	// x*y
//	// x^y
func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for y, _ := range picture {
		for x := 0; x < dx; x++ {
			val := uint8(x * y)
			picture[y] = append(picture[y], val)
		}
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
