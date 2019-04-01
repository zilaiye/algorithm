package main

import (
	"fmt"
)

// 这个是传统的冒泡算法
func BubbleSort(a []int) {
	len := len(a)
	for i := 0 ; i < len  ; i++ {
		for j := i + 1 ; j < len  ; j++ {
			if a[i] > a[j] {
				a[i],a[j] = a[j],a[i]
			}
		}
	}
}

// 这个是网上看到的冒泡算法，但是这个的时间复杂度可能比上面的那个还高，但是能处理最优的情况
func BubbleSort2(a []int) {
	len := len(a)
	swaped := true  
	for swaped {
		swaped = false  
		for i := 0 ;  i < len - 1 ; i++ {
			if a[i] > a[i+1] {
				a[i],a[i+1] = a[i+1],a[i]
				swaped = true 
			}
		}
	}
}

// 冒泡算法的改进算法,往前多看一位数据
func BubbleSort3( a []int ) {
	len := len(a) 
	for i := 0 ; i < len  ; i++ {
		for j := 0 ; j < len - i - 1 ; j++ {
			if a[j] > a[j+1] {
				if j + 2 < len && a[j+1] > a[j+2] {
					a[j] , a[j+2] = a[j+2], a[j]
					j++
				} else {
					a[j] , a[j+1] = a[j+1] , a[j]
				}
			}
		}
	}
	
}

func main() {
	a := []int{1,4,3,2,9,5,6}
	fmt.Printf("origin array :%v \n" , a )
	BubbleSort3(a)
	fmt.Printf("after bubble sort : %v \n",a)
}

