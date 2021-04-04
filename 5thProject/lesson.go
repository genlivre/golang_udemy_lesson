// 文字列型の基礎レッスン

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + "World")

	// 文字列の順序を指定して出力
	fmt.Println("Hello World"[0])         // この書き方だとASCIIが表示される。
	fmt.Println(string("Hello World"[0])) //なので、string方へシフトしてあげる。

	// 指定した順序の文字を置き換える
	var s string = "Hello World"
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	// 特定の文字列が含まれているか
	fmt.Println(strings.Contains(s, "World"))

	// リテラル
	fmt.Println(`Test
											Test
	Test`)

	// ダブルクウォートを表示させたいとき
	// 2パターンあり、
	// 表示させたいダブルクォートの前にバックスラッシュをつける
	fmt.Println("\"")
	// 表示させたいダブルクォートをバッククォートで囲う
	fmt.Println(`"`)
}
