package main

import (
	"fmt"
	"gocode/mygo2/province_from_gov/province"
	"strings"
	"time"
)

func main() {
	baseUrl := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/"
	extraUrl := "index.html"
	proNode := province.NewSearchType1(baseUrl, extraUrl, ".provincetr")
	provinceNode := proNode.Search()

	for _, pN := range provinceNode {
		// 一级地址
		fmt.Println(fmt.Sprintf("%+v",pN))
		// 获取一级地址的 前缀
		provinceHreUrlArr := strings.Split(pN.HrefUrl, ".")
		provincePrefix := provinceHreUrlArr[0]

		//二级地址
		proNode2 := province.NewSearchType2(baseUrl, pN.HrefUrl, ".citytr", pN.AddrCode)
		provinceNode2 := proNode2.Search()
		for _, pN2 := range provinceNode2 {
			fmt.Println(fmt.Sprintf("%+v",pN2))
			// 三级地址
			proNode3 := province.NewSearchType2(baseUrl, pN2.HrefUrl, ".countytr", pN2.AddrCode)
			provinceNode3 := proNode3.Search()

			loop3:
			for _, pN3 := range provinceNode3 {
				if len(pN3.HrefUrl) == 0 { // 三级地址没有
					continue loop3
				}

				fmt.Println(fmt.Sprintf("============================================%+v", pN3))
				time.Sleep(100 * time.Millisecond)
				// 四级地址
				hrefUrl := fmt.Sprintf("%s/%s", provincePrefix, pN3.HrefUrl)
				proNode4 := province.NewSearchType2(baseUrl, hrefUrl, ".towntr", pN3.AddrCode)
				provinceNode4 := proNode4.Search()
				for _, pN4 := range provinceNode4 {
					fmt.Println(fmt.Sprintf("%+v", pN4))
				}
			}
		}
	}
}
