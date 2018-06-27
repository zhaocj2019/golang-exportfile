package export

//ExportCancel 取消导出服务
type ExportCancel struct {
	BaseStruct
	Res bool
}

//Run do action export cancel
func (exportCancel *ExportCancel) Run() bool {
	exportCancel.Res = false
	exportCancel.Res = new(ModelExportOperate).FlagStatus(&ModelExport{Status: STATUSCANCEL})
	return exportCancel.Res
}
