// newとmakeの違い

package main

import "fmt"

func main() {
	// ポインタの中身を指定しないまま、メモリの領域を確保したい場合、newを用いることでｍアドレスが返る
	var p *int = new(int)
	fmt.Println(p)

	// newを記述しないと、nilが返る
	var p2 *int
	fmt.Println(p2)

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>.

	fmt.Println("\n")
	// makeとの違い

	s := make([]int, 0) // スライス
	fmt.Printf("%T\n", s)

	m := make(map[string]int) // マップ
	fmt.Printf("%T\n", m)

	ch := make(chan int) // チャネル
	fmt.Printf("%T\n", ch)

	fmt.Printf("%T\n", p)

	var st = new(struct{}) // ストラクト
	fmt.Printf("%T\n", st)

	// このように、ポインタを返すものはnew、返さないものはmakeを使用する。
}
