package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func add2(a, b int) int {
	return a + b
}

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println(res) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}



/*

语法糖 - 不定长参数 - 默认缺省值

	- func foo(str ...string) {}
	- foo("a", "b")
      
	
闭包 + 匿名函数 + defer

	闭包不是私有啊！闭的意思不是"封闭内部状态"，而是"封闭外部状态"。
	一个函数如何能"封闭外部状态"呢？当外部状态的 scoop 失效的时候，还有一份留在内部状态里面。 -- vazh


*/