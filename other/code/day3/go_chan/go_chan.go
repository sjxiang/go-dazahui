package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		// Fix 2：添加个 loop 变量
		i := i
		go func() {
			fmt.Println(i)  
		}()

		/* Fix 1: 形参
		go func(i int) {
			fmt.Println(i)
		}(i)
		*/
		/* BUG
		go func() {
			fmt.Println(i)  
		}()
		*/
	}
	
	time.Sleep(1 * time.Second)


	shadowExample()

	
	ch := make(chan string)
	
	{  // 异步
		go func() {
			ch <- "hi"  // 发
		}()
		msg := <- ch    // 收
		fmt.Println(msg)	
	}
	

	{  
		go func() {
			for i := 0; i < 3; i++ {
				msg := fmt.Sprintf("message #%d", i+1)
				ch <- msg
			}
			close(ch)
		}()

		for msg := range ch {
			fmt.Println("got:", msg)
		}

		msg := <- ch  // ch 已经 closed
		fmt.Printf("closed: %#v\n", msg)

		msg, ok := <- ch  // ch 已经 closed
		fmt.Printf("closed: %#v (ok=%v)\n", msg, ok)

		// ch <- "hi"  // ch 已经 closed，-> panic
	}


	fmt.Println(sleepSort([]int{32, 12, 4, 15, 27}))
}


/* channel 语义

	操作 - channel 状态 - 结果
	________________________________________	

	send  -  open  - 阻塞，直到 recv
	recv  -  open  - 阻塞，直到 send
	close -  open  - 关闭

	send  -  closed  - panic 
	recv  -  closed  - 不阻塞，获取相对应零值
	close -  closed  - panic

	send  -  nil  - 永久阻塞
	recv  -  nil  - 永久阻塞
	close -  nil  - panic


	详情，https://www.353solutions.com/channel-semantics

	有缓冲的 channel，不会阻塞

*/


// 搞事情
func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, n := range values {
		n := n
		go func() {
			time.Sleep(time.Duration(n) * time.Millisecond)
			ch <- n
		}()
	}

	var out []int
	// for i := 0; i < len(values); i++ {
	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
}


func shadowExample() {
	n := 7
	{                      // `{}` 就是一个 scope，符号表 - 助记符 - 变量名，覆盖 cover
		n := 2
		// n = 2
		fmt.Println("内", n)
	}
	fmt.Println("外", n)
}

/*
	
$ go run -gcflags="-m -l" go_chan.go 

`i escapes to heap`  # 逃逸分析，i 分配在堆上
...

Closure

	闭包不是私有啊！闭的意思不是"封闭内部状态"，而是"封闭外部状态"。
	一个函数如何能"封闭外部状态"呢？
	当外部状态的 scope 失效的时候，还有一份留在内部状态里面。 -- vazh

	e.g. 
		1. defer LIFO，先进后出，延期执行
		2. 起 goroutine 完，离开循环作用域。


*/