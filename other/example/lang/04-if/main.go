package main

import "fmt"

func main() {

	if 7%2 == 0 {                 
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {                 // expr 前可以塞个 statement，但作用域受限
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}


/*

	expr 表达式，有返回值
		e.g.	
			1 == 2，返回 false

	statement 一行代码，无返回值
		e.g. 
			var i int = 1024

*/
