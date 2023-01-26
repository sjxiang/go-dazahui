package main

import (
	"fmt"
	"time"
)

func main() {

	a := 2
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:                          // case 可以有多个值
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}

	t := time.Now()
	switch {                            // 没有 expr，等同于 switch true {}
	case t.Hour() < 12:
		fmt.Println("上午")
	default:
		fmt.Println("下午")
	}
}


/*

	配合接口的类型断言 

*/