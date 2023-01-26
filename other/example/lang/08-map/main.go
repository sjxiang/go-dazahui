package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)           // map[one:1 two:2]
	fmt.Println(len(m))      // 2
	fmt.Println(m["one"])    
	fmt.Println(m["unknow"]) 

	r, ok := m["unknow"]     // 查找
	fmt.Println(r, ok)		 // 0 false

	delete(m, "one")         // 删除

	
	m2 := map[string]int{
		"one": 1, 
		"two": 2,
	}
	
	fmt.Println(m2)
}


/*

创建方式，字面量或者 make

*/