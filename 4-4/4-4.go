package main

// rotate 一次遍历完成元素旋转
import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	rotate(s, 2)
	fmt.Println(s)
}

func rotate(slice []int, n int) {
	n %= len(slice)
	tmp := append(slice, slice[:n]...)
	copy(slice, tmp[n:])
}
