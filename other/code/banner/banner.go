package main

import (
	"fmt"
	"unicode/utf8"
)


func main() {
	
	// 可视化
	banner("Go", 6)
	banner("G😊", 6)  

	s := "G😊"
	fmt.Println("len:", len(s))  // 5 = 1 + 4


	// unicode 字符 ~= rune
	for i, r := range s {
		fmt.Println(i, r)  
		if i == 0 {
			fmt.Printf("%c of type %T\n", r, r)
			// rune (int32)
		}
		fmt.Printf("%x\n", r)
	}

	b := s[0]
	fmt.Printf("%c of type %T\n", b, b)
	// byte (uint8)


	// 格式化输出 	
	x, y := 1, "1"
	fmt.Printf("x=%v, y=%v\n", x, y)
	fmt.Printf("x=%#v, y=%#v\n", x, y)  // 适用于 debug 和 log

	fmt.Printf("%20s\n", s)  // 前面填充 20 字符


	// 测试
	fmt.Println("g", IsPalindrome("g"))
	fmt.Println("go", IsPalindrome("go"))
	fmt.Println("gog", IsPalindrome("gog"))
	fmt.Println("gogo", IsPalindrome("gogo"))

}

// 判断回文
func IsPalindrome(s string) bool {
	rs := []rune(s)
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs)-i-1] {
			return false
		}
	}

	return true
}

// 横幅
func banner(text string, width int) {
	// padding := (width - len(text)) / 2  // BUG: len 是 bytes 数量（字节数组的长度），而不是 char 数量
	padding := (width - utf8.RuneCountInString(text)) / 2 
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

/*

	1. string 延伸

	2. unicode 基础
		unicode table 字符集
			0 => U+0030

		utf-8 传输格式，压缩编码


		
	3. "fmt" 格式化输出


	type stringStruct struct {
		str unsafe.Pointer
		len int
	}


	*/
