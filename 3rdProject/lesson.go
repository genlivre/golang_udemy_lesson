// constの演習
// constは普遍の変数。もう書き換えることのない変数。
package main

import "fmt"

const Pi = 3.14
const (
	UserName = "test_user"
	Password = "test_pass"
)

// varではオーバーフローしてしまう処理であっても、constは頭で型を指定しないので処理ができる。
// var big int = 9223372036854775807 + 1
const big = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi, UserName, Password)
	fmt.Println(big - 1)
}
