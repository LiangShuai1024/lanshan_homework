package main

import (
	"fmt"
	"math/rand"
	"time"
)

func binarySearch(nums [100]int, num int, low int, high int) int {
	for {
		if low > high {
			break
		}
		mid := (low + high) / 2
		if nums[mid] < num {
			low = mid + 1
		} else if nums[mid] > num {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

/*
	for low <= high {
			mid := low + (high-1)/2
			if (*nums)[mid] == num {
				return mid
			} else if (*nums)[mid] > num {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		return -1
	}
*/
func main() {
	var r int
	fmt.Println("现生成一个随机数r...")
	rand.New(rand.NewSource(time.Now().UnixNano()))
	r = rand.Intn(100) + 1
	var a [100]int
	for i := 1; i < 101; i++ {
		a[i-1] = i
	}
	/*for i := 1; i < 101; i++ {
		fmt.Println(a[i-1])
		//fmt.Println(a[i-1])
	}*/
	spot := binarySearch(a, r, 0, 100)
	fmt.Println("经二分查找法得出该数为：", a[spot])
	fmt.Println("而正确答案r为：", r)
}
