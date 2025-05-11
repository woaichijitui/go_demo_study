package _case

import "fmt"

type user struct {
	Name string
	Age  uint
}

func F1() {
	//初始化1
	u := &user{
		Name: "htt",
		Age:  17,
	}
	//初始化2
	//u1 := new(user)

	f2(u)
	fmt.Println(u)

	u2 := new(user)
	u2.Name = "htt2"
	u2.Age = 2
	//f2(u2)
	fmt.Println(u2)
}
func f2(u *user) {

	u.Age = 26
}
