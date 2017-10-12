package main

import "fmt"
import (
	"golang.org/x/tour/wc"
)


type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}



func WordCount(s string) map[string]int {
	
	for w := range(words){
		wrd = words[w]
		if wrd not present in your map:
		add it 
		else 
		increment it 
	}
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
