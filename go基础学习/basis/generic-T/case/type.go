package _case

import "fmt"

type user struct {
	Id   int64
	Name string
	Age  uint8
}

type address struct {
	Id       int
	Province string
	Citi     string
}

// 集合转列表
func mapToList[K comparable, T any](mp map[K]T) []T {
	list := make([]T, len(mp))
	var i = 0
	for _, value := range mp {
		list[i] = value
		i++
	}
	return list
}

// 利用通道打印
func myPrintln[T any](ch chan T) {
	//当你使用 for range 遍历一个 channel 时，循环会从 channel 中依次读取数据，直到 channel 被关闭。如果 channel 未关闭，循环会一直阻塞等待新的数据。
	for data := range ch {
		fmt.Println(data)
	}
}

// 转变为list ，并通过通道c打印
//	泛型使得两种不同类型的map顺利转化

func TTypeCase() {

	u1 := user{
		Id:   1,
		Name: "htt",
		Age:  26,
	}
	u2 := user{
		Id:   2,
		Name: "ljl",
		Age:  23,
	}
	mp1 := map[int]user{1: u1, 2: u2}

	//这里开始看不懂了
	ch := make(chan user)
	//	mp1转为list
	list := mapToList(mp1)

	go myPrintln(ch)
	for _, u := range list {
		ch <- u

	}

	addr1 := address{
		Id:       1,
		Province: "htt",
		Citi:     "河源",
	}
	addr2 := address{
		Id:       2,
		Province: "ljl",
		Citi:     "重庆",
	}
	addMap := map[int64]address{1: addr1, 2: addr2}

	//这里开始看不懂了
	ch1 := make(chan address)
	//	mp1转为list
	addrlist := mapToList(addMap)

	//启用携程
	go myPrintln(ch1)
	for _, addr := range addrlist {
		ch1 <- addr

	}
	close(ch1)
}

// 定义泛型切片
type List[T any] []T

// 定义泛型集合
type mapT[K comparable, T any] map[K]T

// 定义泛型通道
type chanT[T any] chan T

func TTypeCase1() {

	u1 := user{
		Id:   1,
		Name: "htt",
		Age:  26,
	}
	u2 := user{
		Id:   2,
		Name: "ljl",
		Age:  23,
	}
	//定义泛型map后，声明一个自定义泛型map 一定要在其后面加上[] 并限定它的类型
	mp1 := mapT[int, user]{1: u1, 2: u2}

	ch := make(chanT[user])
	//	mp1转为list
	var list List[user]
	list = mapToList[int, user](mp1)

	go myPrintln(ch)
	for _, u := range list {
		ch <- u

	}

	addr1 := address{
		Id:       1,
		Province: "htt",
		Citi:     "河源",
	}
	addr2 := address{
		Id:       2,
		Province: "ljl",
		Citi:     "重庆",
	}
	addMap := mapT[int, address]{1: addr1, 2: addr2}

	//这里开始看不懂了
	ch1 := make(chanT[address])
	//	mp1转为list
	var addrList List[address]
	addrList = mapToList(addMap)

	//启用携程
	go myPrintln(ch1)
	for _, addr := range addrList {
		ch1 <- addr

	}
}
