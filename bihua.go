package bihua

import (
	"strings"
)

// GeNum 获取笔划总数
func GeNum(str string) (num int) {
	// 循环每个字符
	for _, v := range []rune(str) {
		for i, i2 := range data {
			if strings.Contains(i2, string(v)) {
				num += i
			}
		}
	}
	return num
}
