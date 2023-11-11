package main

import "fmt"

type FuncT func(int, int) int

func Cal(a, b int, f FuncT) (result int) {
	result = f(a, b)
	return
}
func Add(a, b int) int {
	return a + b
}
func Minus(a, b int) int {
	return a - b
}
func Times(a, b int) int {
	return a * b
}
func Divide(a, b int) int {
	return a / b
}
func main() {
	fmt.Println("这是一个简单计算器\nAdd >>> +\nMinus >>> -\nTimes >>> *\nDivide >>> /")
	//fmt.Println("Add >>> +" +
	//"Minus >>> -")
	fmt.Println("现在，请输入两个元素,以及您需要的操作：（如：1 1 Add)")
	var num1, num2 int
	var f string
	//var g FuncT
	m := make(map[string]FuncT)
	m["Add"] = Add
	m["Minus"] = Minus
	m["Times"] = Times
	m["Divide"] = Divide
	fmt.Scanf("%d %d %s", &num1, &num2, &f)
	//fmt.Scanf("%d %d ", &num1, &num2)
	//fmt.Println(f)
	fmt.Println("result=", Cal(num1, num2, m[f]))
}
