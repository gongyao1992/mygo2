package aboutdata

import (
	"errors"
	"strings"
)

var DistanceFileConf = map[int]string{
	2 : "scope_city",
	3 : "full_addr",
	4 : "addr_1",
	5 : "addr_2",
	6 : "addr_3",
	7 : "addr_4",
}

type DistanceEle struct {
	ScopeCity 	string `json:"scope_city"`
	FullAddr 	string `json:"full_addr"`
	Addr1 		string `json:"addr_1"`
	Addr2 		string `json:"addr_2"`
	Addr3	 	string `json:"addr_3"`
	Addr4 		string `json:"addr_4"`
	HangInt		string `json:"hang"`
}


// 获取起点
func (d *DistanceEle)GetFromPoint() (*point, error) {
	if len(d.ScopeCity) == 0 {
		return nil, errors.New("港口信息为空，无法处理")
	}
	
	scopeMap := map[string]*point{
		"上海": &point{
			City:     "上海市",
			District: "浦东新区",
			Street:   "顺通路5-9号",
		},
		"深圳": &point{
			City:     "深圳市",
			District: "南山区",
			Street:   "西港路",
		},
		"天津": &point{
			City:     "天津市",
			District: "滨海新区",
			Street:   "跃进路",
		},
	}

	if _, ok := scopeMap[d.ScopeCity]; !ok {
		return nil, errors.New("港口信息超出范围，不需要计算")
	}

	return scopeMap[d.ScopeCity], nil
}
// 获取终点
func (d *DistanceEle)GetToPointArr() ([]*point, error) {
	// 先获取四级地址数量
	addr4Arr := strings.Split(d.Addr4, ",")

	var toArr []*point
	if len(addr4Arr) > 1 {
		toArr = make([]*point, 0)

		// 多个四级地址 可能多个三级地址
		for addrKey, addr4 := range addr4Arr {
			p := point{
				City:     getAddr(d.Addr2, addrKey),
				District: getAddr(d.Addr3, addrKey),
				Street:   addr4,
			}
			toArr = append(toArr, &p)
		}
	} else {
		toArr = make([]*point, 0)
		// 单个四级地址
		// 对详细地址进行处理
		addr := d.Addr1 + d.Addr2 + d.Addr3 + d.Addr4
		detailArr := strings.Split(d.FullAddr, addr)
		for _, detail := range detailArr {
			tDetail := strings.ReplaceAll(detail, " ", "")
			tDetail = strings.ReplaceAll(tDetail, ",", "")
			if len(tDetail) == 0 {
				continue
			}

			p := point{
				City:     d.Addr2,
				District: d.Addr3,
				Street:   d.Addr4,
				Detail: tDetail,
			}
			toArr = append(toArr, &p)
		}
		if len(toArr) == 0 {
			p := point{
				City:     d.Addr2,
				District: d.Addr3,
				Street:   d.Addr4,
			}
			toArr = append(toArr, &p)
		}
	}

	return toArr, nil
}

func getAddr(addrStr string, key int) string {
	addrArr := strings.Split(addrStr, ",")
	addrMap := make(map[int]string)

	if len(addrArr) > 0 {
		for i, addr := range addrArr {
			addrMap[i] = addr
		}
	}

	// 返回数据
	for key >= 0 {
		if _, ok := addrMap[key]; ok {
			return addrMap[key]
		}
		key--
	}

	return ""
}