package export

import "net/http"

//Progress the export Progress struct
type Progress struct {
	Status   int    `json:status`   //status
	ID       string `json:id`       //
	Progress int    `json:progress` //
	Message  string `json:message`  //
}

//ExportProgress export progress struct
type ExportProgress struct {
	id            []string
	RequestParams *map[string]interface{}
}

//SetRequestParams init struct ExportStart
func (ep *ExportProgress) SetRequestParams(requestParams *map[string]interface{}) (t *ExportProgress) {
	ep.RequestParams = requestParams
	return ep
}

//New return  Progress object
func (ep *ExportProgress) New(r *http.Request) *ExportProgress {
	return ep
}

//GetProgress return the  progress status of the export
func (ep *ExportProgress) GetProgress() *[]Progress {
	var res []Progress
	for _, v := range ep.id {
		progress := new(Progress)
		progress.Status = 1
		progress.ID = v
		res = append(res, *progress)
	}
	return &res

}
