package main

import (
	"fmt"
	"strings"
)

func main()  {
	s := "abdaddads"
	fmt.Print(fun(s))
}

/*
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。
 */
func fun(str string) bool {
	if strings.Count(str, "") > 3000 {
		return false
	}

	for _, v := range str {
		fmt.Print(v)
		if v > 127 {
			return false
		}

		if strings.Count(str, string(v)) > 1  {
			return false
		}
	}

	return true
}