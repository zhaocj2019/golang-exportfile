package export

import "net/http"

//ExportCancel 取消导出服务
type ExportCancel struct {
	Res bool
}

//New init an ExportCancel object
func (exportCancel *ExportCancel) New(r *http.Request) *ExportCancel {

	return exportCancel
}

//Run do action export cancel
func (exportCancel *ExportCancel) Run() bool {
	exportCancel.Res = false
	exportCancel.Res = new(ModelExportOperate).FlagStatus(&ModelExport{Status: STATUSCANCEL})
	return exportCancel.Res
}
