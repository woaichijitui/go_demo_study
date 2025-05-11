package _case

import (
	"fmt"
	"sort"
)

//排序

type user struct {
	ID   int
	Name string
	Age  uint
}

func SortCase() {
	var userList []user = []user{
		user{12, "htt", 13},
		user{1, "htt", 45},
		user{22, "htt", 22},
		user{32, "htt", 25},
		user{142, "htt", 27},
		user{125, "htt", 38},
		user{125, "htt", 33},
		user{122, "htt", 99},
		user{122, "htt", 32},
	}
	//	实现各类排序
	sort.Slice(userList, func(i, j int) bool {
		return userList[i].Age > userList[j].Age
	})
	fmt.Println(userList)

	//	实现sort 接口排序
}
