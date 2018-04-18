package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* 这是第一个最简单程序 */
	fmt.Println("hello")
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("输入你的名字：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("出错: %s \n", err)
	} else {
		fmt.Println(input)
		input = input[:len(input)-1]
		fmt.Printf("你好,%s \n", input)
	}
}
