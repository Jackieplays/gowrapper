package main

import "fmt"
import wr "github.com/gowrapper"

func main() {
	list := []int{1, 2, 3, 4}
	in := make(chan wr.Wrap, len(list))
	for x := range list {
		in <- x
	}
	out := wr.Map(map2, in)
	for el := range out {
		fmt.Println(el)
	}
}

func map2(x wr.Wrap) wr.Wrap {
	return x.(int) * 2
}
