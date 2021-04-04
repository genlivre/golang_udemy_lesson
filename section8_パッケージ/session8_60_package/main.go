// package

// コードが始まるファイルはmain.goという命名にすることが多い

package main

import "udemy_lesson/session8_60_package/mylib"
import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
}
