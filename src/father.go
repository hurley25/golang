// Go 组合
package main

import "fmt"

type Father struct {
	Name string
}

func (fath *Father) PrintClassName() {
	fmt.Println("Father")
}

type Child struct {
	Father
	Age int
}

//func (ch *Child) PrintClassName() {
//	fmt.Println("Child")
//}

func main() {
	child := &Child{}
	child.PrintClassName()
}

