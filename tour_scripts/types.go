package main 

import "fmt"

func main() {
	q := []int {2,3,4,5,11,13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int 
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)



	//  slice capacity 
	printSlice(q)

	q = q[:0]
	printSlice(q)

	q = q[:4]
	printSlice(q)

	q = q[2:]
	printSlice(q)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}