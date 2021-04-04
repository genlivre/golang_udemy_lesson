package mylib

import "fmt"

// 変数はキャピタル（頭文字が大文字）にしてあげないと、外部から呼び出すことが出来ない

type Person struct {
	Name string
	Age  int
}

func Say() {
	fmt.Println("Hello")
}
