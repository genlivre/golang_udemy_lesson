// mapレッスン

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)
	m["new"] = 500
	fmt.Println(m)

	// mapの中に存在しないものを取り出そうとすると0が返る
	fmt.Println(m["nothing"])

	// mapはvalueと、値が存在しているかどうかのbooleanも返り値として返してくれる。
	v, ok := m["apple"]
	fmt.Println(v, ok)

	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	// mapの初期化
	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)
}
