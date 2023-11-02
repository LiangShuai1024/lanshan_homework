package main

import "fmt"

/*
func main() {
	var x int
	var p string = "为素数"
	fmt.Scanf("%d", &x)
	for i := 2; i < x; i++ {
		if x%i == 0 {
			p = "不为素数"
			break
		}

	}
	fmt.Println(x, p)
}
*/

func prime(x int) string {
	var p string = "为素数"
	for i := 2; i < x; i++ {
		if x%i == 0 || x < 1 {
			p = "不为素数"
			break
		}
	}
	return p
}
func main() {

	var z int
	//fmt.Println(e(z)
	fmt.Println("请输入一个整数:")
	fmt.Scanf("%d", &z)
	result := prime(z)
	fmt.Println(z, result)

}
