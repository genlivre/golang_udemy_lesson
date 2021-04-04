package main

import (
	"fmt"
	"sort"
)

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d", "a", "f"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)

	// ソートをかける
	sort.Ints(i)
	fmt.Println(i)

	sort.Strings(s)
	fmt.Println(i, s)

	// Nameのアルファベット順で並び替える
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(i, s, p)

	// Ageの若い順で並び替える
	sort.Slice(p, func(i, j int) bool { return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)

}
