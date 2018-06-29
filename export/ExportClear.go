package export

//ExportClear 清除任务
type ExportClear struct {
	BaseStruct
}

//Run 运行
func (ec *ExportClear) Run() bool {
	return true
}
