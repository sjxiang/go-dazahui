package main

import (
	"fmt"
	"math"
)

func main() {

	var a = "init"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"  // 字符串拼接，效率低

	fmt.Println(a, b, c, d, e, f, g)


	const s string = "constant"  // 常量只支持 bool、string、数字
	const h = 500000000          // 没有确定类型的情况下，根据使用的上下文，类型推导
	const i = 3e20 / h

	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}


/*

取变量名：


常量只在编译时，存在


*/
