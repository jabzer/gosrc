package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入你的名字：")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("出错：%s \n", err)
		os.Exit(1)
	} else {
		name := input[:len(input)-1]
		fmt.Printf("你好，%s , 我可以帮你做什么?\n", name)
	}
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("出错：%s \n", err)
			continue
		}
		input = input[:len(input)-1]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			fmt.Println("对不起我帮不了你")
		}

	}

}
