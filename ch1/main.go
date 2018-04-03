package main

import (
	"errors"
	"fmt"
)

var frenchHello string
var emptyString string = ""

func main() {
	fmt.Println("下面是实例：")
	var c complex64 = 5 + 5i
	fmt.Printf("Value is: %v \n", c)
	no, yes, maybe := "no", "yes", "maybe"
	japanseHello := "Konichiwa"
	frenchHello = "Bonjour"
	fmt.Printf("no: %s;yes:%s;maybe:%s", no, yes, maybe)
	fmt.Println(frenchHello, japanseHello)

	//字符串不能变，需要怎样做：
	s := `hello`
	a := []byte(s)
	a[0] = 'c'
	s2 := string(a)
	fmt.Printf("%s \n", s2)
	fmt.Println("\n----------------------------------")
	s1 := "hellos"
	m := " world"
	a1 := s1 + m
	fmt.Printf("%s\n", a1)
	fmt.Println("\n----------------------------------")
	err := errors.New("出错")
	if err != nil {
		fmt.Print(err)
	}

}
