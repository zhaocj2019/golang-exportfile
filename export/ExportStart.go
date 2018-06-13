package export

//@author kalr.zhao<zhaocj2015@163.com>
//@date 2018/06/11
import (
	"encoding/json"
	"errors"

	"github.com/zhaochangjiang/golang-utils/utils"
)

//ExportStart the task export start struct
type ExportStart struct {
	ExportBaseStruct
	Params *ModelExport
}

//Run do action export start function
func (e *ExportStart) Run() (string, error) {

	var id = ""
	err := e.BeforRun()
	if nil != err {
		return id, err
	}
	err = e.paramsDefaultDeal()
	if nil != err {
		return id, err
	}
	id, err = (new(ModelExportOperate)).Add(e.Params)
	if err != nil {
		return id, err
	}

	//缓存任务进度
	progress := new(Progress)
	progress.ID = id
	progress.Message = ""
	progress.Progress = e.Params.Progress
	progress.Status = e.Params.Status
	e.ExportBaseStruct.SaveProgressToCache(progress)

	//执行后台任务生成Excel文件
	go (&RunTask{nil, e.Params.Filename, "", 0, nil, e.Params}).DoTask()
	return id, err
}

//BeforRun 开始执行准备
func (e *ExportStart) BeforRun() error {
	paramJSON, err := json.Marshal(e.RequestParams)
	e.Log.Write("debug", "开始执行导出,参数为:"+string(paramJSON))
	return err
}

//默认参数设置
func (e *ExportStart) paramsDefaultDeal() error {

	var err error
	//如果pageNo 没设置默认为1
	if utils.MapKeyIsSet("pageNo", e.RequestParams) != true {
		(*e.RequestParams)["pageNo"] = 1
	}

	//如果pageSize 没设置，默认指定1000
	if utils.MapKeyIsSet("pageSize", e.RequestParams) != true {
		(*e.RequestParams)["pageSize"] = 1000
	}
	//如果pageSize 没设置，默认指定1000
	if utils.MapKeyIsSet("alpha", e.RequestParams) != true {
		(*e.RequestParams)["alpha"] = 0
	}
	if utils.MapKeyIsSet("module", e.RequestParams) {
		return errors.New("The params module is not exists")
	}

	if utils.MapKeyIsSet("method", e.RequestParams) {
		return errors.New("The params method is not exists")
	}

	if utils.MapKeyIsSet("product", e.RequestParams) {
		return errors.New("The params product is not exists")
	}
	if utils.MapKeyIsSet("filename", e.RequestParams) {
		return errors.New("The params filename is not exists")
	}
	//add data to database
	e.Params = &ModelExport{Status: STATUSSTART, Fileext: "xlsx", Filesize: 0, Createtime: utils.TimeNow(), Progress: 0}

	req := make(map[string]interface{})
	e.Params.Filename, e.Params.RequestModule, e.Params.RequestMethod, e.Params.RequestProduct, e.Params.Alpha, e.Params.Token, e.Params.Generateditself, req = (*e.RequestParams)["filename"].(string), (*e.RequestParams)["module"].(string),
		(*e.RequestParams)["method"].(string),
		(*e.RequestParams)["product"].(string),
		(*e.RequestParams)["alpha"].(int),
		(*e.RequestParams)["token"].(string),
		(*e.RequestParams)["generateditself"].(bool),
		(*e.RequestParams)["param"].(map[string]interface{})

	e.Params.Paramdata = (*e.RequestParams)
	e.Params.RequestFormdata = &req

	//设置开始任务的进度
	e.Params.Progress = PROGRESSSTART
	return err
}

//SetProgress 设置任务进度
func (e *ExportStart) SetProgress() {

}
