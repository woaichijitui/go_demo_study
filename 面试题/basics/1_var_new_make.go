package main

import "fmt"

func VarTest1() {
	type Foo struct {
		name string
		age  int
	}

	//声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %#v\n ", foo1) //main.Foo{age:0, name:""}
	foo1.age = 1
	fmt.Println(foo1.age)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %#v\n ", foo2) //main.Foo{age:0, name:""}
	foo2.age = 2
	fmt.Println(foo2.age)

	//指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %#v\n ", foo3) //&main.Foo{age:0, name:""}
	foo3.age = 3
	fmt.Println(foo3.age)

	//new 初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %#v\n ", foo4) //&main.Foo{age:0, name:""}
	foo4.age = 4
	fmt.Println(foo4.age)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %#v\n ", foo5) //&main.Foo{age:0, name:""}
	foo5.age = 5
	fmt.Println(foo5.age)
}

type StrMap map[string]string

func VarTest2() {
	type Foo struct {
		name string
		age  int
		StrMap
	}

	// 初始化strmap
	// 这样初始化就可以用
	strMap := make(StrMap)
	strMap["http"] = "123"
	strMap["https"] = "456"
	fmt.Println(strMap)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %#v\n ", foo2) //main.Foo{age:0, name:""}
	foo2.StrMap["htt"] = "1234"
	fmt.Println(foo2.StrMap)

	//new 初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %#v\n ", foo4) //&main.Foo{age:0, name:""}
	foo4.StrMap["htt"] = "1234"
	fmt.Println(foo4.StrMap)

	//声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %#v\n ", foo1) //main.Foo{age:0, name:""}
	foo1.StrMap["htt"] = "1234"
	fmt.Println(foo1.StrMap)

	//指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %#v\n ", foo3) //&main.Foo{age:0, name:""}
	foo3.StrMap["htt"] = "1234"
	fmt.Println(foo3.StrMap)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %#v\n ", foo5) //&main.Foo{age:0, name:""}
	foo5.StrMap["htt"] = "1234"
	fmt.Println(foo5.StrMap)
}

// 对比
func VarTest3() {
	type Foo struct {
		name string
		age  int
		StrMap
	}

	//声明初始化
	var foo1 Foo
	fmt.Printf("foo1 --> %#v\n ", foo1) //main.Foo{age:0, name:""}
	foo1.age = 1
	fmt.Println(foo1.StrMap)

	//struct literal 初始化
	foo2 := Foo{}
	fmt.Printf("foo2 --> %#v\n ", foo2) //main.Foo{age:0, name:""}
	foo2.age = 2
	fmt.Println(foo2.StrMap)

	//指针初始化
	foo3 := &Foo{}
	fmt.Printf("foo3 --> %#v\n ", foo3) //&main.Foo{age:0, name:""}
	foo3.age = 3
	fmt.Println(foo3.StrMap)

	//new 初始化
	foo4 := new(Foo)
	fmt.Printf("foo4 --> %#v\n ", foo4) //&main.Foo{age:0, name:""}
	foo4.age = 4
	fmt.Println(foo4.StrMap)

	//声明指针并用 new 初始化
	var foo5 *Foo = new(Foo)
	fmt.Printf("foo5 --> %#v\n ", foo5) //&main.Foo{age:0, name:""}
	foo5.age = 5
	fmt.Println(foo5.StrMap)
}

type Param map[string]interface{}
type Show struct {
	Param
}

// 这是面试题目
func test4() {
	s := new(Show)
	s.Param["RMB"] = 10000
}
