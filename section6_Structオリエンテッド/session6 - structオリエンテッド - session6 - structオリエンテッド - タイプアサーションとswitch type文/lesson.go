// タイプアサーションとswitch type文

package main

import "fmt"

// interface{} を指定すれば、どんな型でも引数として渡せる
func do(i interface{}) {
	ii := i.(int) // interfaceで受け取っただけでは型が決定していないので、使用する際にはこのようにタイプアサーションしてあげる必要がある。
	ii *= 2
	fmt.Println(ii)
}

// stringを受け取ってみる
func doString(i interface{}) {
	ss := i.(string)
	fmt.Println(ss + "!")
}

// でもstringやinteger、booleanなど型関係なく対応できるようにしたいよね？
func doSwitchType(i interface{}) {

	// switch typeというものを使う。
	switch v := i.(type) { // switchとtypeはセット。
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println(v + "!")
	default:
		fmt.Printf("I don't know type %T \n", v)
	}
}

func main() {
	do(10)

	// この様にも書ける。
	var i interface{} = 20
	do(i)

	doString("Mike")

	// switch type
	doSwitchType(100)
	doSwitchType("Ken")
	doSwitchType(true)
}
