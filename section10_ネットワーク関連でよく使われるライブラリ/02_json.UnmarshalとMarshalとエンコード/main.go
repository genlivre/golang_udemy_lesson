package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name,omitempty"`
	Age       int      `json:"age,omitempty"`
	Nicknames []string `json:"nicknames"`
}

// Jsonをカスタマイズする
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name      string   `json:"name,omitempty"`
		Age       int      `json:"age,omitempty"`
		Nicknames []string `json:"nicknames"`
	}{
		Name:      "Mr." + p.Name,
		Age:       p.Age,
		Nicknames: p.Nicknames,
	})
	return v, err
}

func main() {
	// jsonの形式で入ってきたデータをstructで変換する: Unmarshal
	b := []byte(`{"name":"mike", "age":20, "nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	// 今度はjsonの形式に戻す: Marshal
	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
