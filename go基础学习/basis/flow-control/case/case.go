package _case

import "fmt"

func Flow_Control() {
	forCase()
	var user1 user
	switchCase("A", user1)
	switchCase("B", 1)
	switchCase("C", 32.2)
	switchCase("D", 2)

}

// for
func forCase() {
	//for 正常循环
	for i := 0; i < 2; i++ {
		fmt.Printf("for 位置1 打印 %d \n", i)

	}

	//	for 省略条件循环
	var flag = true
	for flag {

		fmt.Printf("for 位置2 打印 %t \n", flag)
		break
	}
	//	for 遍历slice mp string
	var s []int = []int{1, 2, 3, 4, 5}

	for index, data := range s {
		if data == 4 {
			continue
		}
		fmt.Printf("for 位置3 index = %d,data =%d\n", index, data)
	}

	var mp = map[string]string{"A": "a", "B": "b", "C": "c"}

	for key, value := range mp {
		fmt.Printf("for 位置4 key = %s,value =%s\n", key, value)
	}
	var str = "今天天气很好"
	for index, char := range str {
		fmt.Printf("for 位置5 index = %d,char =%c\n", index, char)

	}

	//	for 死循环
	i := 0
	for {
		fmt.Printf("for 位置6 i= %d\n", i)
		i++
		if i == 4 {
			break
		}
	}
	//	for 嵌套循环 break 和 lab：的使用
	fmt.Printf("for 位置7 lab1 , i = %d, \n", i)
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {

			fmt.Printf("for 位置7 i = %d,j =%d\n", i, j)
			if j == 1 {
				return
			}
		}

	}

}

// switch
func switchCase(str string, in interface{}) {

	//	switch 正常
	switch str {
	case "A":
		fmt.Println("情况为A")
	case "B":
		fmt.Println("情况为B")
		fallthrough
	case "C":
		fmt.Println("情况为C")
	default:
		fmt.Println("误匹配字符串")

	}
	i := 0
lab1:
	//	switch 判断类型
	switch in.(type) {
	case user:
		fmt.Println("此类型为user")
	case int:

		fmt.Println("此类型为int")
	default:
		fmt.Println("无匹配类型")

	}

	if i == 0 {

		i++
		goto lab1
	}
	//	goto lab
}

type user struct {
	name string
	age  uint
}
