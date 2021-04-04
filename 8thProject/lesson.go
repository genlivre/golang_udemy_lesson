// 配列とスライスレッスン

package main

import (
	"fmt"
)

func main() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	// 配列は最初に宣言した際のサイズが変動できない。
	// b = append(b, 300) // appendはサイズを変更する関数だが、これは配列ではできない。
	fmt.Println(b)

	// スライスは型のサイズを宣言しないため、サイズの変更ができる。
	var c []int = []int{100, 200}
	c = append(c, 300)
	fmt.Println(c)

	// スライスについて
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])   // 0,1,"2"番目の数字、つまり3が表示される
	fmt.Println(n[2:4]) // 2番目以上4番目未満の数字が表示される
	fmt.Println(n[:4])  // 4番目未満の数字が表示される
	fmt.Println(n[2:])  // 2番目以上の数字が表示される
	fmt.Println(n[:])   // 頭から最後までの数字が表示される

	// 値の書き換え
	n[2] = 100
	fmt.Println(n)

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)

	// スライスの入れ子
	var board = [][]int{
		[]int{0, 1, 2},
		[]int{3, 4, 5},
		[]int{6, 7, 8},
	}
	fmt.Println(board)

}
