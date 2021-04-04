// struct

package main

import "fmt"

type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}
func changeVertex2(v *Vertex) {
	v.X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)

	// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	fmt.Println("\n")

	cv := Vertex{1, 2, "test"}
	changeVertex(cv)
	fmt.Println(cv)

	cv2 := &Vertex{1, 2, "test"}
	changeVertex2(cv2)
	fmt.Println(cv2)
	fmt.Println(*cv2)
}
