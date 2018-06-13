package export

import "net/http"

//Progress the export Progress struct
type Progress struct {
	Status   int8   `json:status`   //status
	ID       string `json:id`       //
	Progress int    `json:progress` //
	Message  string `json:message`  //
}

//ExportProgress export progress struct
type ExportProgress struct {
	ExportBaseStruct
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
