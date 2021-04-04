// rangeレッスン

package main

import "fmt"

func main() {
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	// もっと簡単にかける！

	for i, v := range l {
		fmt.Println(i, v)
	}

	// l（スライス）の中身を取ってきて、インデックス番号と、中身の値を取ってきてくれるのがrange

	// じゃあインデックス番号がいらないときは？　→　_（アンダースコアに置き換える）

	for _, v := range l {
		fmt.Println(v)
	}

	// sliceじゃなく、mapでも同様にできる。

	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// keyだけ取り出す場合は_（アンダースコア）に置き換えなくても良い。
	for k := range m {
		fmt.Println(k)
	}

	// でもvalueだけ取り出したい場合は_（アンダースコア）に置き換える必要がある。
	for _, v := range m {
		fmt.Println(v)
	}

}
