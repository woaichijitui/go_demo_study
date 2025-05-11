package _case

import "errors"

// 形参
// 值传递
func Add(a int, b int) (sum int, err error) {
	if a == 0 && b == 0 {
		err = errors.New("不能同时为零")
		return
	}

	sum = a + b
	return
}

// 引用传递
func Add1(a int, b *int) (sum int, err error) {
	if a == 0 && *b == 0 {
		err = errors.New("不能同时为零")
		return
	}

	sum = a + *b
	*b += 1
	return
}

// 局部变量和全局变量
var G int32 = 40

// 函数中的方法
type user struct {
	name string
	age  uint
}

func NewUser() *user {
	return new(user)
}

func (u *user) GetName() string {
	return u.name

}

func (u *user) GetAge() uint {
	return u.age

}
