package main

import (
	"fmt"
)

var v = "1,2,3"

func main() {
	fmt.Println(v)
	v := []int{1, 2, 3}
	fmt.Println(v)
	if v != nil {
		var v = 123
		fmt.Println(v)

	}

	ints := []int{10, 20, 30, 40}
	for i, d := range ints {
		fmt.Printf("Index: %d,value : %d.     \n", i, d)
	}

    for a,b := range[0:100]{
        fmt.Printf("a:%d,b:%b.\n", a,b)
    }


}
