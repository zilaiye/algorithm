package main 

import (
	"fmt"
)

func QuickSort(a []int , left , right int ) {
	if left < right {
		i , j := left , right 
		value := a[(i+j)/2]

		for i <= j {
			for a[i] < value {
				i++
			}
			for a[j] > value {
				j--
			}

			if i <= j {
				a[i] , a[j] = a[j] , a[i]
				j--
				i++
			}
		}
		if j > left {
			QuickSort(a,left,j)
		}
		if right > i {
			QuickSort(a,i,right)
		}
	}
}

func main() {
	a := []int{10,9,23,14,25,9,1}
	fmt.Println(a)
	QuickSort(a,0,len(a)-1)
	fmt.Println(a)
}

