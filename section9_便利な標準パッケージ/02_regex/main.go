// 正規表現

package main

import (
	"fmt"
	"regexp"
)

func main() {
	/* マッチさせたい文字列が
	頭文字が「a」
	末尾が「e」
	間が「a〜z」で、１つ以上（+）
	という条件で判別する。
	*/
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)

	// 正規表現のプログラムを何度も使いたい場合は、正規表現の部分のみを分離してあげる。
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	// URLの末尾の一致を正規表現か判別する
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)

	// マッチした一部分を取り出すには？
	fss := r2.FindStringSubmatch("/view/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	fss = r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])

	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}
