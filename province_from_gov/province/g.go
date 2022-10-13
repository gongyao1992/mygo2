package province

import "strings"

// AddrEleNode Node 返回的数据节信息
type AddrEleNode struct {
	AddrCode		string // 地址编号
	AddrName 		string // 地址名称
	ParentAddrCode 	string // 父级地址code
	HrefUrl			string // 跳转网址
}

// IAddrNode 地址节点的接口
type IAddrNode interface {
	Search() []*AddrEleNode
}

func dealCode(str string) string {
	// 650100000000
	maxL := 12
	for i := len(str); i < maxL; i++ {
		str += "0"
	}
	return str
}

func cleanStr(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	return str
}