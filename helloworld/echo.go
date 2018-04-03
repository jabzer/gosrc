// echo
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[0:])
	for i := 0; i < len(os.Args[0:]); i++ {
		fmt.Println(i, os.Args[i])
	}

	fmt.Println(strings.Join(os.Args[0:], " "))
}
