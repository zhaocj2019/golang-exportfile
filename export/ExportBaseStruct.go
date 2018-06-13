package export

import (
	"time"

	"github.com/astaxie/beego/cache"
)

//ExportBaseStruct 导出的基础结构体
type ExportBaseStruct struct {
	RequestParams *map[string]interface{}
	Log           *Log
	ProgressID    string
	Cache         cache.Cache
}

//GetProgress 获得任务进度
func (e *ExportBaseStruct) GetProgress() interface{} {
	var v interface{}
	v = e.Cache.Get(e.ProgressID)
	return v
}

//SaveProgressToCache 保存导出进度到缓存中
func (e *ExportBaseStruct) SaveProgressToCache(progressObject *Progress) {
	key := e.GetCacheKey()
	e.Cache.Put(key, *progressObject, 10*time.Minute)

}

//GetCacheKey 获得任务进度缓存key
func (e *ExportBaseStruct) GetCacheKey() string {
	return "export_down_" + e.ProgressID
}
