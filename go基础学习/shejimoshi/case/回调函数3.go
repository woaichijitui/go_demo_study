package huidiao

import "strconv"

// 回调函数 提供实现日志的接口
type Callback3 func(msg string)

// 提供实现
func stringToInt(i int, str string, callback Callback3) string {
	i, err := strconv.ParseInt(str, 10, 64)

}

func Test3() {

}
