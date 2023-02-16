/*
 * @Author: guhang987 610203816@qq.com
 * @Date: 2023-02-16 13:55:02
 * @LastEditors: guhang987 610203816@qq.com
 * @LastEditTime: 2023-02-16 15:46:42
 * @FilePath: /go/hello_go_book/4-7/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

// 翻转utf-8字符元素
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	b := []byte("一 二 三")
	reverseUTF8(b)
	fmt.Println(b)
}

func reverseUTF8(b []byte) {
	for i := 0; i < len(b); {
		// utf-8转10进制unicode
		_, size := utf8.DecodeRune(b[i:])
		reverse(b[i : i+size])
		i += size
	}
}

func reverse(b []byte) {
	last := len(b) - 1
	for i := 0; i < len(b)/2; i++ {
		b[i], b[last-i] = b[last-i], b[i]
	}
}
