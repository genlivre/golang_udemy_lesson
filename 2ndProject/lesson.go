// 変数の演習
package main

import "fmt"

func main() {
	var (
		i      int     = 1
		f64    float64 = 1.2
		f32    float32 = 1.2
		s      string  = "test"
		t      bool    = true
		f      bool    = false
		t2, f2 bool    = true, false
	)
	fmt.Println(i, f64, f32, s, t, f, t2, f2)

	// 変数宣言無しで変数指定をする方法。
	// var宣言は関数外でも使用できる。
	// := ショート宣言は関数内でしか使用できない。
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

	fmt.Printf("%T\n", f64)
	fmt.Printf("%T\n", f32)
}
