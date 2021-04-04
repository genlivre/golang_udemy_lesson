// constに連番を振る

package main

import "fmt"

const (
	c1 = iota
	c2 = iota
	c3 = iota
	c4
	c5
)

const (
	_      = iota
	KB int = 1 << (10 * iota)
	MB
	GB
	TB
)

func main() {
	fmt.Println(c1, c2, c3, c4, c5)
	fmt.Println(KB, MB, GB, TB)
}
