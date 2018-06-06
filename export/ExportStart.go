package export

import (
	"fmt"
	"reflect"
)

//author kalr.zhao<zhaocj2015@sohu.com>
//

//ExportStart the task export start struct
type ExportStart struct {
	Data          []interface{}
	TotalPage     int
	RepsonseBody  string
	fileResource  *ExcelWriter //
	fileName      string
	Params        *ModelExport
	RequestParams interface{}
}

//New init struct ExportStart
func (e *ExportStart) New() (t *ExportStart) {
	return e
}
func (e *ExportStart) paramsDefaultDeal() {
	//add data to database
	e.Params = &ModelExport{Status: STATUSSTART, Filename: e.fileName}
	var types = reflect.TypeOf(e.Params)
	fmt.Println(e.RequestParams)
	fmt.Println(types)
}

//Run do action export start function
func (e *ExportStart) Run() bool {
	e.paramsDefaultDeal()
	(new(ModelExportOperate)).Add(e.Params)

	//Excel handle initialize
	e.fileResource = new(ExcelWriter).New()
	e.fileResource.InitFile().InitSheet(e.fileName)
	defer e.fileResource.Close()

	//get data and deal it
	e.dataDeal()

	return true
}
func (e *ExportStart) dataDeal() {
	//get the first page data and set it at e.Data
	e.firstPage()

	//other page data deal
	e.dealOtherPageData()
}

//dealOtherPageData
func (e *ExportStart) dealOtherPageData() {

	for i := 1; i < e.TotalPage; i++ {
		curl := (new(CURL)).New()
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

//orgReturnData format return Data
func (e *ExportStart) orgReturnData() {

}

//firstPage catch the fist page data
func (e *ExportStart) firstPage() {

	//get data by curl
	e.getData()

	//validate data is right
	e.validateResponseData()

}

//validateData validate the format of the url return data
func (e *ExportStart) validateResponseData() {

}

//GetData fetch the data by curl
func (e *ExportStart) getData() {
	curl := (new(CURL)).New()
	curl.URI = ""
	curl.RequestData.GetParams = map[string]string{}
	curl.RequestData.PostParams = map[string]interface{}{}
	curl.RequestData.Cookies = map[string]string{}
	curl.RequestData.Headers = map[string]string{}
	e.RepsonseBody = curl.Request()
}
