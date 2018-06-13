package export

import (
	"fmt"
	"net/http"

	httprequest "github.com/zhaochangjiang/golang-utils/httprequest"
	"github.com/zhaochangjiang/golang-utils/mycache"
)

const (
	//STATUSSTART 初始化状态
	STATUSSTART = 0

	//STATUSSUCCESS 结束状态
	STATUSSUCCESS = 1

	//STATUSCANCEL 用户取消状态
	STATUSCANCEL = 2

	//STATUSFAILURE 操作失败
	STATUSFAILURE = -1

	//PROGRESSSTART 任务启动时的进度
	PROGRESSSTART = 3

	//PROGRESSGETDATAFINISH 获取数据结束时的进度
	PROGRESSGETDATAFINISH = 96

	//PROGRESSOVER 结束时的进
	PROGRESSOVER = 100
)

//Export 导出服务参数
type Export struct {
	Params     *map[string]interface{}
	LogPointer *Log
	Cache      mycache.Cache
}

//New 初始化对象
func (export *Export) New(r *http.Request) *Export {

	//初始化环境参数
	export.initEnv(r)

	//延迟处理数据
	defer export.Error()
	defer export.LogPointer.Close()
	return export
}

//初始化基础数据
func (export *Export) initEnv(r *http.Request) {

	//初始化 日志
	export.initLog()
	var err error
	export.Cache, err = mycache.NewCache("redis", "{\"password\":\"cae0f7fcf1\",\"conn\":\"172.16.1.39:6379\",\"dbNum\":\"0\"}")
	if nil != err {
		panic(err)
	}
	//初始化参数
	export.Params = new(httprequest.RequestParamsFormat).Run(r)
}

//initLog 初始化日志
func (export *Export) initLog() {

	export.LogPointer = NewLog()
}

//错误信息收集
func (export *Export) Error() {
	if err := recover(); err != nil {
		fmt.Println("RECOVERED:", err)
	}
}

//ExportStart 结构体的方法
func (export *Export) ExportStart() string {

	action := &ExportStart{ExportBaseStruct{RequestParams: export.Params, Log: export.LogPointer, ProgressID: "", Cache: export.Cache}, nil}
	res, err := action.Run()
	if nil != err {
		panic(err)
	}
	return res
}

//ExportCancel 结构体的方法
func (export *Export) ExportCancel() bool {
	action := &ExportCancel{ExportBaseStruct{RequestParams: export.Params, Log: export.LogPointer, ProgressID: "", Cache: export.Cache}, false}
	res := action.Run()
	return res
}

//ExportGetProgress 结构体的方法
func (export *Export) ExportGetProgress() *[]Progress {
	action := &ExportProgress{ExportBaseStruct{RequestParams: export.Params, Log: export.LogPointer, ProgressID: "", Cache: export.Cache}, nil}
	res := action.Run()
	return res
}

//ExportGetList 结构体的方法
func (export *Export) ExportGetList() *[]ExportListReturn {
	action := &ExportGetList{ExportBaseStruct{RequestParams: export.Params, Log: export.LogPointer, ProgressID: "", Cache: export.Cache}, nil}
	res := action.Run()
	return res
}
