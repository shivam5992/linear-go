package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	z := make([][]uint8, dy)
	
	for k := range(z){
		z[k] = make([]uint8, dx)
	}
	
	for i := range(z){
		for j := range(z[i]){
			z[i][j] = uint8(i*j / 2)	
		}
	}
	
	return z
}

func main() {
	pic.Show(Pic)
}