// ポインタ

package main

import "fmt"

func one(x *int) {
	// n = 100が入っていたアドレスを引数xで受け取り、1に書き換える
	*x = 1
}

func main() {
	var n int = 100
	// 普通のint型の数字
	fmt.Println(n)

	// &（アンパサンド）をつけると、この数字（n）が入っているアドレスを表示する
	fmt.Println(&n)

	// アドレスを入れる型には*(アスタリスク)がつく
	var p *int = &n

	fmt.Println(p)

	// アドレスは*（アスタリスク）をつけることで内部を参照することが出来る
	fmt.Println(*p)

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	fmt.Println("\n")

	// デリファレンス
	// アドレスの中身を書き換えること

	// n = 100 のアドレスを関数oneへ渡す
	one(&n)
	fmt.Println(n)
}
