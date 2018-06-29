package export

import (
	"github.com/zhaochangjiang/golang-utils/myutils"
)

const (

	//EveryGetDataPageSize 每次并发取多少页
	EveryGetDataPageSize = 10

	//EveryGetDataSheetSize 每次并发取多少个sheet
	EveryGetDataSheetSize = 10
)

//RunTask 执行生成数据功能操作
type RunTask struct {
	BaseStruct
	FileResource *ExcelWriter //
	FileName     string
	TotalPage    int
	Data         []interface{}
	StartParams  *ModelExport
}

//Run 开始启动任务
func (e *RunTask) Run() {
	//Excel handle initialize
	e.FileResource = NewExcelWriter()
	e.FileResource.InitFile().InitSheet(e.FileName)
	defer e.FileResource.Close()

	//获取数据处理逻辑
	e.dataDeal()
}

//orgReturnData format return Data
func (e *RunTask) orgReturnData() {

}

//获取数据请求
func (e *RunTask) getData() {
	switch e.StartParams.Alpha {
	case 1: //如果是新框架访问
		NewAccessAlpha()
		break
	case 2: //如果是EWS访问
		NewAccessEws()
		break
	default: //如果是老框架访问
		NewAccessZendao()
		break
	}
}
func (e *RunTask) getSheetCount() int {
	return 1
}

//开始数据整合
func (e *RunTask) dataDeal() {

	//获取当前导出功能共多少个sheet
	sheetCount := e.getSheetCount()

	//分页切分，并行获取不同的sheet数据
	sheetMaxPage := myutils.CalcucateMaxPage(sheetCount, EveryGetDataSheetSize)
	for m := 0; m < sheetMaxPage; m++ {

		e.getDataSheet(&getSheetParams{PageSize: e.getSheetTaksCount(sheetCount, m, sheetMaxPage), SheetFromIndex: m * EveryGetDataSheetSize, SheetMax: sheetMaxPage})
	}

}

//responseDataFormat 校验返回数据是否正确,并格式化返回数据用以适配写入Excel文件
func (e *RunTask) responseDataFormat() {

}

type getSheetParams struct {
	PageSize       int
	SheetFromIndex int
	SheetMax       int
}

//获得当前切片改获取的sheet个数
func (e *RunTask) getSheetTaksCount(sheetCount int, formIndex int, sheetMaxPage int) int {
	//获得当前切片需要执行获取任务sheet数量
	var nowPageSize = EveryGetDataSheetSize
	if formIndex+1 == sheetMaxPage {
		sps := sheetCount % EveryGetDataSheetSize
		if sps != 0 { //如果最后一页数量EveryGetDataSheetSize
			nowPageSize = sps
		}
	}
	return nowPageSize
}

//
func (e *RunTask) getDataSheet(sheetParams *getSheetParams) {
	chs := make([]chan int, sheetParams.SheetMax)

	for i := sheetParams.SheetFromIndex; i < sheetParams.SheetFromIndex+sheetParams.PageSize; i++ {
		chs[i] = make(chan int)
		go e.getSheet(chs[i], i)
	}
	for _, ch := range chs {
		<-ch
	}
}

func (e *RunTask) getSheet(ch chan int, sheetIndex int) {
	//get the first page data and set it at e.Data
	//获得首页信息
	e.getData()

	//校验返回数据的格式
	e.responseDataFormat()

	//从第二页开始获取数据
	for i := 2; i < e.TotalPage; i++ {
		e.getData()
		e.responseDataFormat()
	}
	ch <- sheetIndex
}
