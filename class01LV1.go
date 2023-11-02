package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	for {
		var a int
		var b int
		fmt.Println("请输入两个整数，数之间以空格相隔")
		fmt.Scanf("%d %d", &a, &b)
		result := add(a, b)
		fmt.Println("它们的和为：", result)
	}
}
