package province

import (
	"fmt"
	"testing"
)

func TestNewSearchType2(t *testing.T) {
	//"http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/2021/14/01/140105.html"
	baseUrl := "http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/"
	extraUrl := "01/140105.html"
	extraUrl = "2021/14/01/140105.html"

	proNode3 := NewSearchType2(baseUrl, extraUrl, ".towntr", "140100000000")
	provinceNode3 := proNode3.Search()
	for _, pN3 := range provinceNode3 {
		fmt.Println(fmt.Sprintf("%+v", pN3))
	}
}