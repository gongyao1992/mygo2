package province

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"time"
)

type addrSearchType2 struct {
	baseUrl			string
	extraUrl		string
	filterStr 		string
	parentCode		string
}

func NewSearchType2(baseUrl, extraUrl, filterS, parentCode string) *addrSearchType2 {
	return &addrSearchType2{
		baseUrl: baseUrl,
		extraUrl: extraUrl,
		filterStr: filterS,
		parentCode: parentCode,
	}
}

func (p *addrSearchType2)Search() []*AddrEleNode {
	time.Sleep(100 * time.Millisecond)
	// 获取页面信息
	url := p.baseUrl + p.extraUrl
	// TODO 动态代理 ip
	//client := http.Client{
	//	Transport:     &http.Transport{
	//		Proxy: http.ProxyURL(uri),
	//	},
	//} https://yr6dkm0rue.feishu.cn/docs/doccnWNxID61Y2hhe7cmAXiS3Pc
	
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Println("http error:")
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
	doc.Find(p.filterStr).Each(func(i int, select1 *goquery.Selection) {
		n := new(AddrEleNode)
		select1.Find("td").Each(func(i int, select3 *goquery.Selection) {
			text := cleanStr(select3.Text())
			if i == 0 {
				n.AddrCode = text
			} else if i == 1 {
				n.AddrName = text
			}

			hrefSelect := select3.Find("a")
			n.HrefUrl = hrefSelect.AttrOr("href", "")
			// 父节点
			n.ParentAddrCode = p.parentCode
		})
		ret = append(ret, n)
	})

	return ret
}
