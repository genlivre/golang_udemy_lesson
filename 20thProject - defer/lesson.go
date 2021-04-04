// deferレッスン

package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("world foo")

	fmt.Println("Hello foo")
}

func stackingDefer() {
	// deferは下から順に実行される。

	fmt.Println("run")
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println("success")
}

func fileOpen() {
	// じゃあどういうときにdeferって使うの？
	file, _ := os.Open("./lesson.go") // ファイルをオープンする
	defer file.Close()                // その直後にdeferで必須処理であるCloseを書いておけば、忘れない。
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}

func main() {
	foo()
	stackingDefer()
	fileOpen()

	// deferは、関数内の実行が終了した後に実行される。
	defer fmt.Println("world")

	fmt.Println("Hello")

}
