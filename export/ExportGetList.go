package export

//ExportListReturn every export return of the export
//@author karl.zhao<zhaocj2009@126.com>
//@date 2018/06/05
type ExportListReturn struct {
	FileName string
	Status   int
	Progress int
}

//ExportGetList get export list
type ExportGetList struct {
	DataListFromDataBase *[]ModelExport
	RequestParams        *map[string]interface{}
}

//SetRequestParams init struct ExportGetList
func (exportGetList *ExportGetList) SetRequestParams(requestParams *map[string]interface{}) (t *ExportGetList) {
	exportGetList.RequestParams = requestParams
	return exportGetList
}

//GetIDList get base message from database
func (exportGetList *ExportGetList) GetIDList() {
	exportGetList.DataListFromDataBase = (new(ModelExportOperate)).FetchList()
}

//GetList get the export message
func (exportGetList *ExportGetList) GetList() *[]ExportListReturn {
	var res []ExportListReturn
	exportGetList.GetIDList()
	for _, v := range *exportGetList.DataListFromDataBase {
		exportListReturn := new(ExportListReturn)
		exportListReturn.FileName = v.Filename
		res = append(res, *exportListReturn)
	}
	return &res
}
