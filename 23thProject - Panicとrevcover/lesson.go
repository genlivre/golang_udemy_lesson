// panicとrecover

package main

import "fmt"

func thirdPartyConnectDB() {
	panic("Unable to connect database!")
	// Goはpanicを自分のコードの中で書くことはあまり推奨していない
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}

func main() {
	save()
	fmt.Println("OK?")
}
