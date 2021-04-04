// 演習

package main

import "fmt"

func q1() {
	// Q1 以下のスライスから一番小さい数を探して出力するコードを書いてください。
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	var minNum int
	for i, num := range l {
		if i == 0 {
			minNum = num
			continue
		}
		if minNum >= num {
			minNum = num
		}
	}
	fmt.Println(minNum)
}

func q2() {
	// 以下の果物の価格の合計を出力するコードを書いてください。
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var sum int
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("Q1の答えは")
	q1()
	fmt.Println("Q2の答えは")
	q2()
}
