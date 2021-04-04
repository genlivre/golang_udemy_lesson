// メソッドとポインタレシーバーと値レシーバー

package main

import "fmt"

type Vertex struct {
	X, Y int
}

func Area(v Vertex) int {
	return v.X * v.Y
}

// メソッド
func (v Vertex) Area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

// ポインターレシーバー

func main() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))

	// メソッドは.（ドット）で呼び出せる
	fmt.Println(v.Area())

	// ポインターレシーバーでVertexの値を書き換える
	v.Scale(10)
	fmt.Println(v.Area())
}
