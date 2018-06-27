package export

import (
	"strconv"
)

//RunTask 执行生成数据功能操作
type RunTask struct {
	BaseStruct
	FileResource *ExcelWriter //
	FileName     string
	RepsonseBody string
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

	//get data and deal it
	e.dataDeal()
}

//orgReturnData format return Data
func (e *RunTask) orgReturnData() {

}
func (e *RunTask) getDataAlpha() {
	crd := &CurlRequestData{}
	e.callCurl(crd)
}

func (e *RunTask) getDataEws() {
	crd := &CurlRequestData{}
	e.callCurl(crd)
}
func (e *RunTask) getOldFrameWork() {
	crd := &CurlRequestData{}
	e.callCurl(crd)
}

//发送请求获取数据
func (e *RunTask) callCurl(crd *CurlRequestData) {

	var err error
	var time = 3

	curl := NewCurl()
	curl.URI = crd.URI
	curl.RequestData.GetParams = crd.GetParams
	curl.RequestData.PostParams = crd.PostParams
	curl.RequestData.Cookies = crd.Cookies
	curl.RequestData.Headers = crd.Headers

	//未防止一次请求数据超时，发送3次请求
	for i := 0; i < time; i++ {
		e.RepsonseBody, err = curl.Request()
		if nil == err { //如果有一次获取数据成功，那么就表示成功,退出循环和方法
			return
		}
		e.Log.Write(LogError, "第"+strconv.Itoa(i)+"次尝试请求数据失败,内容描述"+err.Error())
	}
	e.Log.Write(LogError, "连续尝试"+strconv.Itoa(time)+"次获取数据失败!")
}

//获取数据请求
func (e *RunTask) getData() {
	switch e.StartParams.Alpha {
	case 1: //如果是新框架访问
		e.getDataAlpha()
		break
	case 2: //如果是EWS访问
		e.getDataEws()
		break
	default: //如果是老框架访问
		e.getOldFrameWork()
		break
	}
}
func (e *RunTask) getSheetCount() int {

	return 1
}
func (e *RunTask) getSheet(ch chan bool, sheetIndex int) {
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
	ch <- true
}
func (e *RunTask) dataDeal() {

	sheetCount := e.getSheetCount()
	chs := make([]chan bool, sheetCount)
	for i := 0; i < sheetCount; i++ {
		chs[i] = make(chan bool)
		go e.getSheet(chs[i], i)
	}
	for _, ch := range chs {
		<-ch
	}
	e.FileResource.AddRow()
}

//responseDataFormat 校验返回数据是否正确,并格式化返回数据用以适配写入Excel文件
func (e *RunTask) responseDataFormat() {

}
