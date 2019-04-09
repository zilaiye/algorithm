package main
 
import (
	// "strings"
	"fmt"
)
 
/*
	给出一个字符串，求出不重复的子字符串的最长长度
*/
 
func lengthOfLongestSubstring(s string) int {
	length := 0
	len := len(s) 
	if len > 0 {
		length = 1
	}
	ss := []byte(s)
	for i := 0 ; i < len ; i++ {
		flag := 0 
		for j := i + 1 ; j < len ; j++ {
			for k := i ; k < j ; k++ {
				if ss[k] == ss[j] {
					flag = 1 
					break
				}
			}
			if flag > 0 {
				break
			}
			if j - i + 1 > length {
				length = j - i + 1 
			}
		}
	}
    return length
}

func main() {
	s1 := "abcabcbb"
	fmt.Println("s1 : ", lengthOfLongestSubstring(s1))
	s2 := "bbbbb"
	fmt.Println("s2 : " , lengthOfLongestSubstring(s2))
	s3 := "pwwkew"
	fmt.Println("s3 : " , lengthOfLongestSubstring(s3))
}