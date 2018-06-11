package export

import (
	"errors"
	"net/http"
	"strconv"
)

//ModelExport 数据结构体
type ModelExport struct {
	Filename        string `json:"filename"`       //文件名称
	Fileext         string `json:"fileext"`        //文件扩展名
	Filesize        int64  `json:"filesize"`       //文件大小
	Status          int8   `json:"status"`         //导出状态
	Createtime      string `json:"createtime"`     //创建时间
	RequestProduct  string `json:"requestproduct"` //获取数据的渠道
	RequestModule   string //获取数据的模块
	RequestMethod   string //获取数据的方法
	Paramdata       interface{}
	RequestHeader   *http.Header            //请求对方服务器的Header信息
	RequestFormdata *map[string]interface{} //获取数据传参
	Userinfo        *Userinfo               //用户信息
	Generateditself bool                    //是否自动生成
	Progress        int                     // 导出的进度
	Alpha           int
	Token           string
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
