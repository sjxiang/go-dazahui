package main

import "fmt"

func add2(n int) {
	n += 2
}

func add2ptr(n *int) {
	*n += 2
}

func main() {
	
	n := 5
	
	add2(n)
	fmt.Println(n) // 5

	add2ptr(&n)
	fmt.Println(n) // 7
}

/*

    &<x> 
	获取变量的地址
    
	*<x>  
	对内存空间的读写

*/