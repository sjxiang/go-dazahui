package main

import (
	"fmt"
	"sort"
)


func main() {
	var s []int  
	fmt.Println("len", len(s))  
	if s == nil {                // slice 零值为 nil
		fmt.Println("nil slice")
	}	

	s2 := []int{1, 2, 3, 4, 5, 6 , 7}
	fmt.Printf("s2 = %#v\n", s2)


	// 切片操作
	s3 := s2[1:4]  // 区间，左开右闭
	fmt.Printf("s2 = %#v\n", s3)

	// fmt.Println(s2[4:100])  // 越界 panic


	s3 = append(s3, 100)  // 副作用，随之而来
	fmt.Printf("s3 (append) = %#v\n", s3)   
	fmt.Printf("s2 (append) = %#v\n", s2)
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))


	var s4 []int
	// s4 := make([]int, 0, 1_000)
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Println("s4", len(s4), cap(s4))

	// 测试 concat
	fmt.Println(concat([]string{"A", "B"}, []string{"C", "D"}))  // [A B C D]

	
	// 测试 median
	vs := []float64{4, 2, 1, 3}
	fmt.Println(median(vs))
	fmt.Println(vs)

	fmt.Println(median(nil))
}


// 中位数
func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("median of empty slice")
	}

	// copy 是为了不改变 values
	nums := make([]float64, len(values))
	copy(nums, values)
	
	sort.Float64s(nums)
	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	v := (nums[i-1] + nums[i]) / 2
	return v, nil
}


// 连接
func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s

	// return append(s1, s2...)
}

// 追加
func appendInt(s []int, v int) []int {
	i := len(s)

	if len(s) < cap(s) {  // 空间充足
		s = s[:len(s)+1]
	} else {  			  // 需要重新开辟新数组，拷贝
		fmt.Printf("重新开辟：%d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s
}


/*

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}



s1 := make([]int, 10)
array, len=10, cap=10
|
|_____________________________________________
                  |                          |
				  |                          |
underlying array [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
indices           0  1  2  3  4  5  6  7  8  9 
						   |        |
        ___________________|________|      
		|
		|
	array, len=4 (7-3), cap=7 (10-3)

s2 := s1[3:7]


- 创建新的切片会复用
- 也会有副作用

*/