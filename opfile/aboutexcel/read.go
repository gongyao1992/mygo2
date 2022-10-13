package aboutexcel

import (
	"fmt"
	excelize "github.com/xuri/excelize/v2"
	"strconv"
)

//func ReadFile(fileDir string) {
//	// 打开文件
//	f, err := excelize.OpenFile(fileDir)
//	if err != nil {
//		return
//	}
//	// 获取全部的数据
//	rows, err := f.GetRows("Sheet1")
//	if err != nil {
//		return
//	}
//	for _, values := range rows {
//
//		for i, value := range values {
//			fmt.Println(i, value)
//		}
//	}
//
//	f.SetCellValue()
//
//	f.SaveAs()
//}

type MyFile struct {
	f *excelize.File
	workingSheet string
}

func OpenFile(dir string) *MyFile {
	f, err := excelize.OpenFile(dir)
	if err != nil {
		panic("dir文件打开失败：" + err.Error())
	}
	return &MyFile{
		f: f,
	}
}

func (myF *MyFile)SetSheet(sheetName string) {
	myF.workingSheet = sheetName
}

func (myF *MyFile)ReadFile(confMap map[int]string) interface{} {
	rows, err := myF.f.GetRows(myF.workingSheet)
	if err != nil {
		panic("sheet错误：" + err.Error())
	}

	formatMapArr := make([]map[string]string, 0, len(rows))

	for hangK, values := range rows {

		formatMap := make(map[string]string)

		// 一行数据
		for i, value := range values {
			if key, ok := confMap[i]; ok {
				formatMap[key] = value
			}
		}
		formatMap["hang"] = strconv.Itoa(hangK + 1)
		formatMapArr = append(formatMapArr, formatMap)
	}

	return formatMapArr
}

func (myF *MyFile)SetValue(hangI, lieStr string, value interface{})  {
	axls := fmt.Sprintf("%s%s", lieStr, hangI)
	myF.f.SetCellValue(myF.workingSheet, axls, value)
}

func (myF *MyFile)Save(outFileName string)  {
	myF.f.SaveAs(outFileName)
}