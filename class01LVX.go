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
	spot := binarySearch(a, r, 1, 101)
	fmt.Println("经二分查找法得出的数为：", a[spot])
	fmt.Println("答案为：", r)
}
