package main 

import "fmt"
import "math"

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func Sqrt(x float64) float64 {
	z := float64(1)
	
	i := 0 
	for i < 10 {
		z_new := z - (z*z - x) / (2 * z)
		i+=1 
		z = z_new 
	}
	return z
}

func main() {
	defer fmt.Println("world")

	sum := 0
	for i:=0 ; i <= 10 ; i ++ {
		sum += i
	}
	fmt.Println(sum)


	new_sum := 1 
	for ; new_sum < 1000; {
		new_sum += new_sum
	}
	fmt.Println(new_sum)


	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)


	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(3))

	fmt.Println("hello")

	
}