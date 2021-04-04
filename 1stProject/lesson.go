package main

import (
	"fmt"
	"os/user"
	"time"
)

/*
func init() { // Initはmain関数内で呼び出しをしなくても、最初に呼ばれる関数
	fmt.Println("Init!")
}
*/

func main() {
	fmt.Println("Hello World", time.Now())
	fmt.Println(user.Current())
}
