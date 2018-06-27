package export

import "net/http"

const (
	//ProgressStepStart 初始导出状态
	ProgressStepStart = 0

	//ProgressStepStartParams 参数格式化完成
	ProgressStepStartParams = 3

	//ProgressTaskStart 导出开始任务状态,准备获取第一页数据
	ProgressTaskStart = 4

	//ProgressTaskFirstPage 获取第一页数据成功状态
	ProgressTaskFirstPage = 5

	//ProgressTaskFinishGetData 获取数据结束状态
	ProgressTaskFinishGetData = 95

	//ProgressTaskCreateFileFinish 生成文件成功状态
	ProgressTaskCreateFileFinish = 96

	//ProgressTaskUploadFileFinish 上传文件完成状态
	ProgressTaskUploadFileFinish = 99

	//ProgressExportFinish 导出完成状态
	ProgressExportFinish = 100
)

//Progress the export Progress struct
type Progress struct {
	Status   int8   `json:status`   //status
	ID       string `json:id`       //
	Progress int    `json:progress` //
	Message  string `json:message`  //
}

//ExportProgress export progress struct
type ExportProgress struct {
	BaseStruct
	ID []string
}

//New return  Progress object
func (ep *ExportProgress) New(r *http.Request) *ExportProgress {
	return ep
}

//Run return the  progress status of the export
func (ep *ExportProgress) Run() *[]Progress {
	var res []Progress
	for _, v := range ep.ID {
		progress := new(Progress)
		progress.Status = 1
		progress.ID = v
		res = append(res, *progress)
	}
	return &res

}
