package main 

import (
	"fmt"
	"math/rand"
	"math"
	"math/cmplx"
)

func add(x int, y int) int{
	return x + y
}

// Arguments Shortned
func add1(x, y int) int{
	return x + y
}

// multiple returns
func swap (x, y string) (string, string){
	return y,x 
}

// naked return 
func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 -1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
		Big = 1 << 100
		Small = Big >> 99
	)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}


func main(){
	// Print Statement
	fmt.Println("My fav number is", rand.Intn(100))

	// Note: Exported names start with Caps
	fmt.Println(math.Pi)

	//  Function Call
	fmt.Println(add(13,31))

	// Multiple Return Function
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(18))


	var i int
	fmt.Println(i, c, python, java)

	// with initialization and dynamic init
	var c, python bool = true, true
	var k = "without datatype"
	fmt.Println(c, python, k)

	// declarations
	l := 10
	a1, b1, c1 := true, false, "no!"
	fmt.Println(a1, b1, c1, l)


	//  Cmplex types 
	fmt.Printf("%T : %v \n", ToBe, ToBe)
	fmt.Printf("%T : %v \n", MaxInt, MaxInt)
	fmt.Printf("%T : %v \n", z, z)

	// Type conversions 
	var x, y int = 3, 4 
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y , z)

	const Pi = 3.14
	// Pi := 34
	fmt.Println(Pi)


	// Big Ints 
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

