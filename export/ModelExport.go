package export

import (
	"net/http"
	"net/url"
)

//ModelExport 数据结构体
type ModelExport struct {
	Filename        string       `json:"filename"`       //文件名称
	Fileext         string       `json:"fileext"`        //文件扩展名
	Filesize        int64        `json:"filesize"`       //文件大小
	Status          int8         `json:"status"`         //导出状态
	Createtime      string       `json:"createtime"`     //创建时间
	RequestProduct  string       `json:"requestproduct"` //获取数据的渠道
	RequestModule   string       //获取数据的模块
	RequestMethod   string       //获取数据的方法
	RequestHeader   *http.Header //请求对方服务器的Header信息
	RequestFormdata *url.Values  //获取数据传参
	Userinfo        *Userinfo    //用户信息
	Generateditself bool         //是否自动生成
	Progress        int          // 导出的进度
}

//ModelExportOperate the operate
type ModelExportOperate struct {
}

//Add
func (m *ModelExportOperate) Add(me *ModelExport) bool {

	return true
}

//FetchList
func (m *ModelExportOperate) FetchList() *[]ModelExport {
	var res []ModelExport
	return &res
}

//FlagStatus
func (m *ModelExportOperate) FlagStatus(modelExport *ModelExport) bool {
	return true
}
