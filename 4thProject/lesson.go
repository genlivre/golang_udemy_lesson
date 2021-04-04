// 数値型の基礎レッスン

package main

import "fmt"

func main() {
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	fmt.Printf("type=%T value=%v\n", u8, u8)

	// Printfについて
	// %T は type（型）を表示してくれる
	// %t は boolean型のtrue or false　を返してくれる
	// integer型は %d で10進数で表示してくれる
	// stringは %s で文字列で返してくれる。
	// などがある。

	// floatで下２桁まで表示する場合は　%.2fなど。

	fmt.Println("【計算演習】")
	fmt.Println(" 1 + 1 =", 1+1)
	fmt.Println(" 10 - 1 =", 10-1)
	fmt.Println(" 10 / 2 =", 10/2)
	fmt.Println(" 10 / 3 =", 10/3)
	fmt.Println(" 10.0 / 3 =", 10.0/3)
	fmt.Println(" 10 / 3.0 =", 10/3.0)
	fmt.Println(" 10 % 2 =", 10%2)
	fmt.Println(" 10 % 3 =", 10%3)

	fmt.Println("【シフト演習】")
	fmt.Println(1 << 0) // 0001 0001
	fmt.Println(1 << 1) // 0001 0010
	fmt.Println(1 << 2) // 0001 0100
	fmt.Println(1 << 3) // 0001 1000
}
