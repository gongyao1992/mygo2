package main

import (
	"fmt"
	"github.com/gongyao1992/go-util/helper"
	"gocode/mygo2/opfile/aboutdata"
	"gocode/mygo2/opfile/aboutexcel"
	"os"
	"sync"
)

func main()  {

	getwd, err := os.Getwd()
	if err != nil {
		return
	}
	fileDir := getwd + "/opfile/21年10月-22年9月各地址箱量.xlsx"
	fmt.Println(fileDir)
	//fileDir := ""

	excel := aboutexcel.OpenFile(fileDir)
	excel.SetSheet("Sheet1")
	i := excel.ReadFile(aboutdata.DistanceFileConf)

	//fmt.Println(i)
	var distanceArr []aboutdata.DistanceEle
	helper.InterfaceToStruct(i, &distanceArr)

	pool := NewPool(50)
	for _, dis := range distanceArr {
		pool.Add(1)
		go func(excel *aboutexcel.MyFile, dis aboutdata.DistanceEle, group *Pool) {
			defer group.Done()
			op(excel, dis)
		}(excel, dis, pool)
	}
	pool.Wait()

	excel.Save("21年10月-22年9月各地址箱量1.xlsx")
}

func op(excel *aboutexcel.MyFile, dis aboutdata.DistanceEle)  {
	fromPoint, err := dis.GetFromPoint()
	fmt.Println(dis)
	if err == nil {
		toPointArr, _ := dis.GetToPointArr()
		pointStruct := aboutdata.NewPointList(fromPoint, toPointArr)
		disSum, err := pointStruct.Deal()
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		}
		excel.SetValue(dis.HangInt, "J", disSum)
		excel.SetValue(dis.HangInt, "K", errMsg)
	} else {
		excel.SetValue(dis.HangInt, "J", "")
		excel.SetValue(dis.HangInt, "K", "起点不参与运算")
	}
}


// Pool Goroutine Pool
type Pool struct {
	queue chan int
	wg *sync.WaitGroup
}
// New 新建一个协程池
func NewPool(size int) *Pool{
	if size <=0{
		size = 1
	}
	return &Pool{
		queue:make(chan int,size),
		wg:&sync.WaitGroup{},
	}
}
// Add 新增一个执行
func (p *Pool)Add(delta int){
	// delta为正数就添加
	for i :=0;i<delta;i++{
		p.queue <-1
	}
	// delta为负数就减少
	for i:=0;i>delta;i--{
		<-p.queue
	}
	p.wg.Add(delta)
}
// Done 执行完成减一
func (p *Pool) Done(){
	<-p.queue
	p.wg.Done()
}
// Wait 等待Goroutine执行完毕
func (p *Pool) Wait(){
	p.wg.Wait()
}