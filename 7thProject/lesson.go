// 型変換レッスン

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	// Atoiメソッドは返り値として数字とエラー文を出す。
	// i, _ := strconv.Atoi(s)
	// 上記のようにerrorメッセージを返さない場合は、_（アンダースコア）を記述することで、そこの処理を宣言しなくても使用できる。
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("ERROR")
	}
	fmt.Printf("%T %v\n", i, i)

	h := "Hello World"
	fmt.Println(string(h[0]))
}
