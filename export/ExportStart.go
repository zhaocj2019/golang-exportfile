package export

//@author kalr.zhao<zhaocj2015@163.com>
//@date 2018/06/11
import (
	"encoding/json"
	"errors"
	"fmt"

	utils "github.com/zhaochangjiang/golang-utils/myutils"
)

//ExportStart the task export start struct
type ExportStart struct {
	BaseStruct
	Params *ModelExport
	Port   int
	Domain string
}

//Run do action export start function
func (e *ExportStart) Run() (string, error) {

	var id = ""
	err := e.BeforRun()
	if nil != err {
		return id, err
	}

	//参数格式化
	err = e.paramsDefaultDeal()
	if nil != err {
		return id, err
	}

	//设置初始的导出进度
	e.setInitProgress()

	//往数据库写入一条导出任务数据
	e.BaseStruct.ProgressID, err = (new(ModelExportOperate)).Add(e.Params)
	if nil != err {
		return id, err
	}
	e.validatePermit()

	//设置初始的导出进度
	e.setInitProgress()

	//执行后台任务生成Excel文件
	(&RunTaskClient{StartParams: e.Params, Domain: e.Domain, Port: e.Port}).Run()
	return e.BaseStruct.ProgressID, err
}

//权限验证
func (e *ExportStart) validatePermit() {
	var version = "v1"
	//  $timer = ['connection_timeout' => 25, 'execute_timeout' => 25];
	api := "ucenter-v1.auth.validate"

	var params = make(map[string]string)

	params["_TOKEN"] = e.Params.Token

	params["url"] = "/" + version + "/" + e.Params.RequestProduct + "/" + e.Params.RequestModule + "/" + e.Params.RequestMethod
	params["subsystem"] = (e.Params.RequestProduct + "-" + version)
	content, err := NewAccessAlpha().SetMethod("POST").SetAPI(api).SetParams(&params).Run()
	if nil != err {
		panic(err)
	}
	fmt.Println(content)
}

//保存进度到Redis
func (e *ExportStart) setInitProgress() {
	//缓存任务进度
	progress := new(Progress)
	progress.ID = e.BaseStruct.ProgressID
	progress.Message = ""
	progress.Progress = e.Params.Progress
	progress.Status = e.Params.Status
	e.BaseStruct.SaveProgressToCache(progress)
}

//BeforRun 开始执行准备
//return error
func (e *ExportStart) BeforRun() error {

	paramJSON, err := json.Marshal(e.RequestParams)

	if err != nil {
		return err
	}
	e.Log.Write(LogDebug, "开始执行导出,参数为:"+string(paramJSON))
	return err
}

//默认参数设置
func (e *ExportStart) paramsDefaultDeal() error {

	var err error
	//从第一页获取数据
	(*e.RequestParams)["pageNo"] = 1

	//如果pageSize 没设置，默认指定1000
	if utils.MapKeyIsSet("pageSize", e.RequestParams) != true {
		(*e.RequestParams)["pageSize"] = 1000
	}
	//如果pageSize 没设置，默认指定1000
	if utils.MapKeyIsSet("alpha", e.RequestParams) != true {
		(*e.RequestParams)["alpha"] = 0
	}

	//校验以下参数是否存在
	for _, v := range [...]string{"module", "method", "product", "filename"} {
		if utils.MapKeyIsSet(v, e.RequestParams) {
			return errors.New("The params " + v + " is not exists")
		}
	}

	//add data to database
	e.Params = &ModelExport{Status: STATUSSTART, Fileext: "xlsx", Filesize: 0, Createtime: utils.TimeNow(), Progress: ProgressStepStart}

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
	e.Params.Progress = ProgressStepStartParams
	return err
}

//SetProgress 设置任务进度
func (e *ExportStart) SetProgress() {

}
