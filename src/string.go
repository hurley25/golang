// string.go

package main

import "fmt"

func main() {
	str1 := "hello"
	str2 := "hello"

	if str1 == str2 {
		fmt.Println("==")
	} else {
		fmt.Println("!=")
	}
}
