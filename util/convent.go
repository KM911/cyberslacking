package util

import "strconv"

func String2Int(str string) int {
	value, _ := strconv.Atoi(str)
	return value
}

// 包裹一个函数 其返回值是 interface{} 类型 err
// 如果不为nil 则 panic
func WarpErr(fc func() (interface{}, error)) interface{} {
	value, err := fc()
	if err != nil {
		panic(err)
	}
	return value
}
