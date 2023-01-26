package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println(sum) // 9

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println(k, v) // b 8; a A
	}
	for k := range m {
		fmt.Println("key", k) // key a; key b
	}
}


/*

range 语法糖，搭配挺多

1. for i, v := range string {}，每次迭代的都是 rune 字符 
	
	e.g. 'e'、"中"

	
	range 字符串遍历
		会调用 runtime/utf8.go 中的 decoderrune 函数，把字符串解码成 rune 类型的字符（例，判断前 3 个字节是中文）
		

		runtime/utf8.go

		type rune int32
		func decoderune(s string, k int) (r rune, pos int) {}


2. for range array {}，循环次数为 len(array)


3. for i, v := array {}，拷贝了一份副本，用来遍历

	等同于，
		tmp := array
		v := tmp[i]

           
*/