package main

import (
	"fmt"
)

func SelectSort(a []int) {
	len := len(a) 
	for i := 0 ; i < len ; i++ {
		min := i 
		for j := i + 1 ; j < len ; j++ {
			if a[j] < a[min] {
				min = j 
			}
		}
		a[i] , a[min] = a[min] ,a[i]
	}
}

func main() {
	a := []int{10,9,23,14,25,9,1}
	fmt.Println(a)
	SelectSort(a)
	fmt.Println(a)
}