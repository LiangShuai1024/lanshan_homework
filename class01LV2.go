package main

import (
	"fmt"
	"math"
)

func o(x float64) float64 {
	return x * x * math.Pi
}
func main() {
	var r float64
	fmt.Println("请输入圆的半径：")
	//fmt.Scanf("%g", &r)
	fmt.Scanf("%g", &r)
	result := o(r)
	fmt.Println("圆的面积约为", result)

}
