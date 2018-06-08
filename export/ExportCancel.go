package export

//ExportCancel 取消导出服务
type ExportCancel struct {
	Res           bool
	RequestParams *map[string]interface{}
}

//SetRequestParams init struct ExportCancel
func (exportCancel *ExportCancel) SetRequestParams(requestParams *map[string]interface{}) (t *ExportCancel) {
	exportCancel.RequestParams = requestParams
	return exportCancel
}

//Run do action export cancel
func (exportCancel *ExportCancel) Run() bool {
	exportCancel.Res = false
	exportCancel.Res = new(ModelExportOperate).FlagStatus(&ModelExport{Status: STATUSCANCEL})
	return exportCancel.Res
}
