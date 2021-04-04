// switchレッスン

package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "mac"
}

func main() {
	os := getOsName()
	switch os {
	case "mac":
		fmt.Println("お使いのOCはMACです")
	case "windows":
		fmt.Println("お使いのOCはWINDOWSです")
	default:
		fmt.Println("DEFAULT", os)
	}

	t := time.Now()
	fmt.Println(t.Hour())
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}
}
