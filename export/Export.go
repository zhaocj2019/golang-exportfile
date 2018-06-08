package export

import (
	"fmt"
	"net/http"

	httprequest "github.com/zhaochangjiang/golang-utils/httprequest"
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
	RequestParams *map[string]interface{}
}

//New 初始化对象
func (export *Export) New(r *http.Request) *Export {
	export.RequestParams = new(httprequest.RequestParamsFormat).Run(r)

	//延迟处理数据
	defer export.Error()
	return export
}

//错误信息收集
func (export *Export) Error() {
	if err := recover(); err != nil {
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	}
}

//ExportStart 结构体的方法
func (export *Export) ExportStart() bool {
	res := new(ExportStart).SetRequestParams(export.RequestParams).Run()
	return res
}

//ExportCancel 结构体的方法
func (export *Export) ExportCancel() bool {
	res := new(ExportCancel).SetRequestParams(export.RequestParams).Run()
	return res
}

//ExportGetProgress 结构体的方法
func (export *Export) ExportGetProgress() *[]Progress {
	res := new(ExportProgress).SetRequestParams(export.RequestParams).GetProgress()
	return res
}

//ExportGetList 结构体的方法
func (export *Export) ExportGetList() *[]ExportListReturn {
	res := new(ExportGetList).SetRequestParams(export.RequestParams).GetList()
	return res
}
