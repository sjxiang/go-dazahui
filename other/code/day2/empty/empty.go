package main

import "fmt"


func main() {

	// go =  1.17
	var i interface{}

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	s := i.(string)  // 类型断言
	fmt.Println("s:", s)

	/*
		n := i.(int)  // panic 
		fmt.Println("n:", n)
	*/

	n, ok := i.(int)
	if ok {
		fmt.Println("i:", n)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) {  // type switch，搭配 interface 的类型断言
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}
	
	
	fmt.Println(maxInts([]int{3, 2, 1}))
}


func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}



