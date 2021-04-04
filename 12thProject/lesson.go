// 関数のレッスン

package main

import "fmt"

func add(x int, y int) int { // 関数の（）の中に引数を記述し、返り値は（）の外側に型を記述する。
	// fmt.Println("add function")
	return x + y
}

func add2(x2, y2 int) (int, int) { // 返り値が複数ある場合は、（）で括り、返り値の数だけ型を記述する。
	return x2 + y2, x2 - y2
}

func cal(price, item int) (result int) { // ここの段階で返り値のresultというint変数を宣言している。
	result = price * item // なので、ここでは「result := 」という表記はしない。
	return                // ネイキッドリターンといい、最初の関数宣言時にリターンする変数名を既に宣言しているので、return後に変数名を指定しなくても、ここでは「result」が返される。
}

func main() {
	r := add(10, 20)
	fmt.Println(r)

	r1, r2 := add2(10, 20)
	fmt.Println(r1, r2)

	r3 := cal(100, 2)
	fmt.Println(r3)

	// 変数に関数を入れる場合
	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(1)

	// 変数に入れずに、その場で走らせることも可能
	func(x int) {
		fmt.Println("inner func", x)
	}(200)
}
