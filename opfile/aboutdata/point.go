package aboutdata

import (
	"errors"
	"fmt"
	"github.com/gongyao1992/go-util/helper"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type point struct {
	City		string `json:"city"`
	District	string `json:"district"`
	Street 		string `json:"street"`
	Detail		string `json:"detail"`
}

type PointEle struct {
	m	map[int]*point
}

func NewPointList(fromP *point, toPArr []*point) *PointEle {

	outP := make(map[int]*point)
	i := 0
	outP[i] = fromP
	i += 1
	for _, toP := range toPArr {
		tP := point{
			City:     toP.City,
			District: toP.District,
			Street:   toP.Street,
			Detail:   toP.Detail,
		}
		outP[i] = &tP
		i += 1
	}

	outP[i] = fromP

	return &PointEle{m: outP}
}

func (pl *PointEle)Deal() (int, error) {

	sumDis := 0
	maxI := len(pl.m)
	i := 0
	for (i + 1) < maxI {
		a := pl.m[i]
		b := pl.m[i + 1]

		disInt, err := getDistance(a, b)
		if err != nil {
			return 0, err
		}

		sumDis += disInt
		i += 1
	}
	return sumDis, nil
}

func getDistance(a, b *point) (int, error) {

	//url := "https://svip.56hello.com/g/dp/bmap/distance/pair2"
	url := "https://svip.56hello.com/g/dp/bmap/distance/piar"

	body := struct {
		From	*point `json:"from"`
		To		*point `json:"to"`
	}{
		From: a,
		To: b,
	}

	type retStruct struct {
		Data struct{
			Distance int `json:"distance"`
			From2To int `json:"from_2_to"`
			To2From int `json:"to_2_from"`
		} `json:"data"`
		ErrorCode int `json:"error_code"`
		ErrorMsg string `json:"error_msg"`
	}

	postI, _ := POST(url, body)
	var retS retStruct
	err := helper.StrToStruct(string(postI), &retS)
	if err != nil {
		fmt.Println(string(postI))
		return 0, err
	}
	if retS.ErrorCode > 0 {
		return 0, errors.New(retS.ErrorMsg)
	}

	return retS.Data.From2To, nil
}

const httpTimoutSecond = time.Duration(30) * time.Second

// POST 发送post请求
func POST(url string, body interface{}) ([]byte, error) {
	// 创建请求
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(helper.ToJson(body)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resultByte, nil
}
