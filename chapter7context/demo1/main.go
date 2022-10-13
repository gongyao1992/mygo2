package main

import (
	"context"
	"fmt"
)

func main()  {
	ccc()

	//var ctx context.Context
	//ctx = context.TODO()
	//ctx1, canel := context.WithCancel(ctx)
	//ctx = context.WithValue(ctx, "key1", "0001")
	//ctx = context.WithValue(ctx, "key2", "0001")
	//ctx = context.WithValue(ctx, "key3", "0001")
	//ctx = context.WithValue(ctx, "key4", "0004")
	//
	//fmt.Println(ctx.Value("key2"))
	//canel()
	//a(ctx1)
	//b(ctx)
}

func a(ctx context.Context)  {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:

		b(ctx)
		ctx.Deadline()
		fmt.Println("--")
	}
}

func b(ctx context.Context)  {
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	default:
		fmt.Println("--")
	}
}

func ccc() {
	gen := func(ctx context.Context) chan<- int { // 通道是分方向的 不能乱用
		ch := make(chan int)
		n := 1
		go func(ch chan int, n int) {
			for {
				select {
				case <- ctx.Done():
					break
				default:
					ch <- n
					n += 1
				}
			}
		}(ch, n)
		return ch
	}

	ctx := context.Background()
	ctxCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := range gen(ctxCancel) {
		fmt.Println(i)
		if i == 6 {
			break
		}
	}
}