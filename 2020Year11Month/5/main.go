package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {
	s := "巩 尧 gong yao"
	fmt.Println(replaceBlank(s))
}

// 请编写一个方法，
// 将字符串中的空格全部替换为“%20”。
// 假定该字符串有足够的空间存放新增的字符，
// 并且知道字符串的真实长度(小于等于1000)，
// 同时保证字符串由【大小写的英文字母组成】。
// 给定一个string为原始的串，返回替换后的string。
func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}

	for _, v := range s {
		fmt.Println(v, " ", string(v), " ", unicode.IsLetter(v))
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}
	}

	return strings.Replace(s, " ", "%20", -1), true
}