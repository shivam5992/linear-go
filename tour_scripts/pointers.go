package main 

import "fmt"


type Vertex struct {
	X int 
	Y int
}

type Vertex1 struct {
	X, Y int
}

func main() {
	i, j := 42, 2701 

	p := &i 
	fmt.Println(*p)
	*p = 21 
	fmt.Println(i)


	p = &j
	*p = *p / 37 
	fmt.Println(j)


	// fmt.Println(Vertex{1,2})

	v := Vertex{1, 2}
	// v.X = 4
	// fmt.Println(v.X)

	p1 := &v
	p1.X = 1e9
	fmt.Println(v)

	var (
	v1 = Vertex1{1, 2}
	v2 = Vertex1{X: 1}
	v3 = Vertex1{}
	p2 = &Vertex1{1, 2}
	)

	fmt.Println(v1, p2, v2, v3)

	var ar [2]string
	ar[0] = "Hello"
	ar[1] = "World"
	fmt.Println(ar[0], ar[1])
	fmt.Println(ar)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)


	var s []int = primes[1:4]
	fmt.Println(s)
}