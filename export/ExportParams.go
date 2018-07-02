package export

//NewEveryExportRequest 请求参数数据初始化
func NewEveryExportRequest() *EveryExportRequest {
	return &EveryExportRequest{}
}

//EveryExportRequest 每个请求的参数
type EveryExportRequest struct {
	Filename       string              `json:"filename"`
	Exportid       string              `json:"exportid"`
	Alpha          int                 `json:"alpha"`
	DownloadParams *[]EverySheetParams `json:"download_params"`
}

//EverySheetParams 每个页签的参数
type EverySheetParams struct {
	Product      string              `json:"product"`
	Module       string              `json:"module"`
	Method       string              `json:"method"`
	RemoteParams []map[string]string `json:"remote_params"`
}
