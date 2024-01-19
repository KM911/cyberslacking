package ds

import (
	"fmt"

	"github.com/boljen/go-bitmap"
)

// 用你来表示文件chunk的收发状态
func TrueValue(bm *bitmap.Bitmap) []int {

	lens := bm.Len()
	result := []int{}
	for i := 0; i < lens; i++ {
		if bm.Get(i) {
			fmt.Println("value i", i, "is true")
			result = append(result, i)
		}
	}
	return result
}

func FalseValue(bm *bitmap.Bitmap) []int {
	lens := bm.Len()
	result := []int{}
	for i := 0; i < lens; i++ {
		if !bm.Get(i) {
			result = append(result, i)
		}
	}
	return result
}

// 不行了啊,感觉必须通过第一个byte来进行类型的判断,主要是每次都是不同的链接
