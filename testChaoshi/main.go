//package main
//
//import (
//	"context"
//	"fmt"
//	"runtime"
//	"time"
//)
//
//const timeOut = 6 * time.Second
//const timeOut2 = 8 * time.Second
//
//// 测试超时
//func main()  {
//	go test()
//	time.Sleep(7 * time.Second)
//	printGoroutine("结束后")
//}
//
//func test()  {
//	ctx := context.Background()
//
//	var cancel context.CancelFunc
//	ctx, cancel = context.WithTimeout(ctx, timeOut)
//	defer cancel()
//
//	c := make(chan bool, 1)
//	go func(c chan bool) {
//		runtime.Goexit()
//		// 内部如何关闭
//		runTime()
//		// SQL
//		//runtime.Goexit()
//		c <- true
//	}(c)
//
//	printGoroutine("结束前")
//
//	select {
//	case <-c:
//		fmt.Println("正常执行结束")
//	case <-ctx.Done():
//		fmt.Println("超时执行结束")
//		runtime.Goexit()
//	}
//
//	return
//}
//
//func runTime()  {
//	time.Sleep(timeOut2)
//}
//
//func printGoroutine(str string)  {
//	n := runtime.NumGoroutine()
//	fmt.Println(fmt.Sprintf("%s 协程数量: %d", str, n))
//}
package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int) // 无缓冲的channel
	go unbufferChan(ch)

	for i := 0; i < 10; i++ {
		//number := <-ch
		//fmt.Println("receive ", number) // 读出值
		select {
		case number := <-ch:
			fmt.Println("receive ", number)
		}
	}
}

func unbufferChan(ch chan int) {
	for i := 0; i < 10; i++ {
		//fmt.Println("send ", i)
		//ch <- i // 写入值
		//time.Sleep(100 * time.Microsecond)
		select {
		case ch <- i:
			fmt.Println("send ", i)
		}
	}
}