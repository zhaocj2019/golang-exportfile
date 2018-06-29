package export

import (
	"errors"
	"net/http"
	"strconv"
)

//ModelExport 数据结构体
type ModelExport struct {
	Filename        string                  `json:"filename,omitempty"`       //文件名称
	Fileext         string                  `json:"fileext,omitempty"`        //文件扩展名
	Filesize        int64                   `json:"filesize,omitempty"`       //文件大小
	Status          int8                    `json:"status,omitempty"`         //导出状态
	Createtime      string                  `json:"createtime,omitempty"`     //创建时间
	RequestProduct  string                  `json:"requestproduct,omitempty"` //获取数据的渠道
	RequestModule   string                  `json:"request_module,omitempty"` //获取数据的模块
	RequestMethod   string                  `json:"request_method,omitempty"` //获取数据的方法
	Paramdata       interface{}             `json:"paramdata,omitempty"`
	RequestHeader   *http.Header            `json:"request_header,omitempty"`   //请求对方服务器的Header信息
	RequestFormdata *map[string]interface{} `json:"request_formdata,omitempty"` //获取数据传参
	Userinfo        *Userinfo               `json:"userinfo,omitempty"`         //用户信息
	Generateditself bool                    `json:"generateditself,omitempty"`  //是否自动生成
	Progress        int                     `json:"progress,omitempty"`         // 导出的进度
	Alpha           int                     `json:"alpha,omitempty"`
	Token           string                  `json:"token,omitempty"`
}

//ModelExportOperate the operate
type ModelExportOperate struct {
}

//Add 添加一条数据
func (m *ModelExportOperate) Add(me *ModelExport) (string, error) {
	var id = 10
	err := errors.New("sdfs")
	return strconv.Itoa(id), err
}

//FetchList 查询进度
func (m *ModelExportOperate) FetchList() *[]ModelExport {
	var res []ModelExport
	return &res
}

//FlagStatus 标记数据状态
func (m *ModelExportOperate) FlagStatus(modelExport *ModelExport) bool {
	return true
}
