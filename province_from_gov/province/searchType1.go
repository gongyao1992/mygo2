package province

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"time"
)

type addrSearchType1 struct {
	baseUrl			string
	extraUrl		string
	filterStr 		string
}

func NewSearchType1(baseUrl, extraUrl, filterS string) *addrSearchType1 {
	return &addrSearchType1{
		baseUrl: baseUrl,
		extraUrl: extraUrl,
		filterStr: filterS,
	}
}

func (p *addrSearchType1)Search() []*AddrEleNode {
	time.Sleep(100 * time.Millisecond)

	// 获取页面信息
	url := p.baseUrl + p.extraUrl
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("http error")
		return nil
	}

	//// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	ret := make([]*AddrEleNode, 0)

	// 第一层进行过滤
	doc.Find(p.filterStr).Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(i int, selection *goquery.Selection) {
			hrefUrl := selection.AttrOr("href", "")
			strArr := strings.Split(hrefUrl, ".")
			n := AddrEleNode{
				AddrCode:       dealCode(strArr[0]),
				AddrName:       selection.Text(),
				ParentAddrCode: "",
				HrefUrl:        hrefUrl,
			}
			ret = append(ret, &n)
		})
	})

	return ret
}