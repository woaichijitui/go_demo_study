package main

import (
	"fmt"
	_simple "gobasis/reflect/1_simple"
)

func main() {

	_simple.Reftype(5)

	mp := map[int]string{1: "name"}
	name := "nihao"
	User := struct {
		Name string
	}{Name: "htt"}

	_simple.Reftype(name)
	_simple.Reftype(User)

	fmt.Println("______________________________________________")
	//2
	_simple.RefValue(5)
	_simple.RefValue(name)
	_simple.RefValue(User)
	_simple.RefValue(mp)

	fmt.Println("______________________________________________")
	//要传指针
	_simple.RefSetValue(&name)
	fmt.Println(name)

	fmt.Println("______________________________________________")
	_simple.GetStructTag()
}
