package export

type RunTask struct {
	fileResource *ExcelWriter //
	fileName     string
	RepsonseBody string
	TotalPage    int
	Data         []interface{}
	StartParams  *ModelExport
}

//DoTask 开始启动任务
func (e *RunTask) DoTask() {

	//Excel handle initialize
	e.fileResource = new(ExcelWriter).New()
	e.fileResource.InitFile().InitSheet(e.fileName)

	defer e.fileResource.Close()

	//get data and deal it
	e.dataDeal()
}

//orgReturnData format return Data
func (e *RunTask) orgReturnData() {

}

//GetData fetch the data by curl
func (e *RunTask) getData() {
	curl := (new(CURL)).New()
	curl.URI = ""
	curl.RequestData.GetParams = map[string]string{}
	curl.RequestData.PostParams = map[string]interface{}{}
	curl.RequestData.Cookies = map[string]string{}
	curl.RequestData.Headers = map[string]string{}
	e.RepsonseBody = curl.Request()
}

//firstPage catch the fist page data
func (e *RunTask) firstPage() {

	//get data by curl
	e.getData()

	//validate data is right
	e.validateResponseData()

}
func (e *RunTask) dataDeal() {
	//get the first page data and set it at e.Data
	e.firstPage()
	//other page data deal
	e.otherPageData()
}

//validateData validate the format of the url return data
func (e *RunTask) validateResponseData() {

}

//dealOtherPageData
func (e *RunTask) otherPageData() {
	curl := (new(CURL)).New()
	for i := 1; i < e.TotalPage; i++ {
		curl.URI = ""
		curl.RequestData.GetParams = map[string]string{}
		curl.RequestData.PostParams = map[string]interface{}{}
		curl.RequestData.Cookies = map[string]string{}
		curl.RequestData.Headers = map[string]string{}
		curl.Request()
		//get data by curl
		e.getData()
		e.validateResponseData()
	}
}
