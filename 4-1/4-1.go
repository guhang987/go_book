// 编写一个函数，用于统计SHA256散列中不同的位数
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	c3 := []byte{'A', '2', '1'}
	c4 := []byte{'A', '2', 1}
	fmt.Printf("%08b\n%08b\n", c3, c4)
	fmt.Println(diff(c3, c4))
	fmt.Printf("%x\n%x\n", c1, c2)
	fmt.Println(diff(c1[:], c2[:]))
}

func diff(b1, b2 []byte) (count int) {
	l1, l2 := len(b1), len(b2)
	var maxL int
	if l1 > l2 {
		maxL = l1
	} else {
		maxL = l2
	}
	// 按字节遍历，统计每个字节中异位个数
	for i := 0; i < maxL; i++ {
		switch {
		case i > l1:
			count += popCount(b2[i])
		case i > l2:
			count += popCount(b1[i])
		default:
			count += popCount(b1[i] ^ b2[i])
		}
	}
	return
}

func popCount(b byte) (count int) {
	for b != 0 {
		b &= b - 1
		count++
	}
	return count
}
