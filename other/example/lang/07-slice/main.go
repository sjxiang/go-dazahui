package main

import "fmt"

func main() {

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s[2])   // c
	fmt.Println("len:", len(s)) // 3

	s = append(s, "d")   
	s = append(s, "e", "f")
	fmt.Println(s) // [a b c d e f]


	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c) 

	fmt.Println(s[2:5]) // [c d e]
	fmt.Println(s[:5])  // [a b c d e]
	fmt.Println(s[2:])  // [c d e f]


	good := []string{"g", "o", "o", "d"}
	fmt.Println(good) // [g o o d]
}

/*

Slice 描述比较贴切，对数组的引用

	- array，数据起始地址
	- len，字节长度；越界 panic
	- cap，容量


创建切片
	- make，runtime.makeSlice()
	- 字面量，先创建数据（array），再塞进实例字段中
	- 根据已有的切片，再次切片


*/


/*

slice
    
    addr unsafe.Pointer  起始的指针 
    len 目前可读写值的长度；panic，越界
    cap 容量

    nums := make([]int, 5)
    // 默认，cap = len

    var nums []int  
    // nil, 0, 0

    

    向其中添加数值
        如果知道 len 与 cap
            nums[i] = val

        不知道
            nums = append(nums, val)
            // tmp := nums 拷贝副本
            // length := len(nums)
            // tmp[lengtn] = val
            // nums = tmp

            nums_1 := append(nums, val)
            // ❌ 别瞎搞 append，原有底层数组引用没有释放，造成内存泄漏
        
            // 扩容，根据 len 与 cap 关系，决定


    根据已有的，制作切片
 
        底层共用一个支撑数组，有副作用；任意一方扩容，就瓦解了

        ```go
        s1 := []string{"0","1", "2", "3", "4", "5"} // len=5 cap=5
        s2 := s1[2:4] // len= 2 cap=4
        // s2 := s1[2:4:4] // len=2 cap=2 提供下标，写死容量了. append 后，原底层数组分开；append 前还是受影响的

        s1[2] = "CHANGED"
        fmt.Println(s1)
        fmt.Println(s2)

        output:
        [0 1 CHANGED 3 4 5]
        [CHANGED 3]

        ```

        手动拷贝，防范副作用

            ```go
            s3 := make([]int, len(s1))
            copy(s3, s1)
            ```









*/


/*
# Go-string


1. 原理

    字符    编号     字节序列
    A       65       0100 0001
    世      19990    1110 0100 1011 1000 1001 0110
    string  rune     byte
    
    步骤：
        收录字符，指定编号，存储数值 => 字符集 => 大一统字符集 Unicode
    
        减少开销，采用不定长编码 UTF-8


2. string

    type string struct {
        addr  // 起始地址 找得到开头，找不到结尾？
        len   // 字节个数 
    }
    

3. 用法

    var str string
    // nil, 0

    str := "pig"

    fmt.Println(str[0])  // 可读
    str[0] = 'b'  // ❌ 字符串内容是不会被修改的，分配到只读内存段
    
    str = "hello" 

    bs := ([]byte)(str) // 类型转换，byte 切片
    bs[0] = 'b'

    
    for i, v := str {  // 迭代以 字符 为单位，也就是 rune
        fmt.Printf("%q\n", str[i])    
    }
    
    计数
        按字符而不是字节
        import "unicode/utf8"
        utf8.RuneCountInString()
*/