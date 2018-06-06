package export

import (
	"net/http"
)

//ExportInterface 导出要实现的方法规范
type ExportInterface interface {
	ExportStart()
	ExportCancel()
	ExportGetProgress()
	ExportGetList()
}

const (
	//STATUSSTART 初始化状态
	STATUSSTART = 0

	//STATUSSUCCESS 结束状态
	STATUSSUCCESS = 1

	//STATUSCANCEL 用户取消状态
	STATUSCANCEL = 2

	//STATUSFAILURE 操作失败
	STATUSFAILURE = 3
)

//Export 导出服务参数
type Export struct {
	ExportAbstract
}

//ExportStart 结构体的方法
func (export *Export) ExportStart(r *http.Request) bool {
	return new(ExportStart).New().Run()
}

//ExportCancel 结构体的方法
func (export *Export) ExportCancel(r *http.Request) bool {
	return new(ExportCancel).New(r).Run()
}

//ExportGetProgress 结构体的方法
func (export *Export) ExportGetProgress(r *http.Request) *[]Progress {
	return new(ExportProgress).New(r).Get()
}

//ExportGetList 结构体的方法
func (export *Export) ExportGetList(r *http.Request) *[]ExportListReturn {
	return new(ExportGetList).New(r).GetList()
}
