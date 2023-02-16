package main

// 用指针，就地反转整个数组
import (
	"fmt"
)

func main() {
	s := [5]int{5, 6, 7, 8, 9}
	remove(&s)
	fmt.Println(s)
}

func remove(slice *[5]int) {
	for i, j := 0, len(*slice)-1; i < j; {
		slice[i], slice[j] = slice[j], slice[i]
		i, j = i+1, j-1
	}

}
