// PublicとPrivate

package main

import "udemy_lesson/session8_61_package/mylib"
import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(mylib.Average(s))

	mylib.Say()
	person := mylib.Person{Name: "Mike", Age: 20}
	fmt.Println(person)
}

