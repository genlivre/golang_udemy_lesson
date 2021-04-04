// 演習問題

package main

import "fmt"

func q1() {
	fmt.Println("Q1. 以下の1.11をint型に変換して出力してください。")
	f := 1.11
	i := int(f)

	fmt.Println(i)
	fmt.Printf("%T %v\n", i, i)
}

func q2() {
	fmt.Println("\nQ2. コードを書かずに以下の出力結果を答えてください。\ns := []int{1, 2, 5, 6, 2, 3, 1}\nfmt.Println(s[2:4])")
	fmt.Println("A.[5 6]")

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[2:4])
}

func q3() {
	fmt.Println("\nQ3. 以下のコードを実行した時に\nfmt.Printf(\"%T %v\", m, m)\n以下のような出力結果となるmを作成してください。\nmap[string]int map[Mike:20 Nancy:24 Messi:30]")

	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}

	fmt.Printf("%T %v", m, m)
}

func main() {
	q1()
	q2()
	q3()
}
