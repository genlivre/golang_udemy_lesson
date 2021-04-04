// go routine が長くかかり過ぎたら途中でキャンセルするcontext

package main

import (
	"context"
	"fmt"
	"time"
)

func longProcess(chx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)

	// 1秒経っても処理が終了しなかったらcancelを走らせる
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// もしキャンセル処理の定義がまだ未決定の場合、TODOを使ってエラーを回避させる
	// ctx := context.TODO()

	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("######################")
}
