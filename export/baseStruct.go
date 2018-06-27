package export

import (
	"time"

	"github.com/astaxie/beego/cache"
)

//BaseStruct 导出的基础结构体
type BaseStruct struct {
	RequestParams *map[string]interface{}
	Log           *Log
	ProgressID    string
	Cache         cache.Cache
}

//GetProgress 获得任务进度
func (e *BaseStruct) GetProgress() interface{} {
	var v interface{}
	v = e.Cache.Get(e.ProgressID)
	return v
}

//SaveErrorMessage 保存错误信息到数据库
func (e *BaseStruct) SaveErrorMessage(mesage string, id string) bool {

	return true
}

//SaveProgressToCache 保存导出进度到缓存中，当前默认保存到Redis中
func (e *BaseStruct) SaveProgressToCache(progressObject *Progress) {

	key := e.GetCacheKey()
	e.Cache.Put(key, *progressObject, 10*time.Minute)

	//如果提示信息不为空 则保存信息
	if progressObject.Message != "" {

		//将错误信息保存到数据库中
		e.SaveErrorMessage(progressObject.Message, progressObject.ID)
	}
}

//GetCacheKey 获得任务进度缓存key
func (e *BaseStruct) GetCacheKey() string {
	return "export_down_" + e.ProgressID
}
