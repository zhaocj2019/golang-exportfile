package export

import (
	"sync"

	curl "github.com/mikemintang/go-curl"
)

//CurlRequestData  the params of the request to server
type CurlRequestData struct {
	URI        string
	Headers    map[string]string
	Cookies    map[string]string
	PostParams map[string]interface{}
	GetParams  map[string]string
}

//CurlResponse curl return contents
type CurlResponse struct {
	err      error
	response *curl.Response
}

//CURLReturnData the return of the Request data struct
type CURLReturnData struct {
	Pagesize    int
	TotalCount  int
	TotalPage   int
	Data        []interface{}
	Err         error
	TableHeader map[string]string
}

type CURL struct {
	requestResponse CurlResponse
	curlReturnData  *CURLReturnData
	RequestData     *CurlRequestData
	URI             string
}

var cr *CURL
var once sync.Once

//New
func (cr *CURL) New() *CURL {
	once.Do(func() {
		cr = &CURL{}
	})
	return cr
}

//Request
func (cr *CURL) Request() string {
	var res = ""
	// 链式操作
	req := curl.NewRequest()
	cr.requestResponse.response, cr.requestResponse.err = req.
		SetUrl(cr.RequestData.URI).
		SetHeaders(cr.RequestData.Headers).
		SetCookies(cr.RequestData.Cookies).
		SetQueries(cr.RequestData.GetParams).
		SetPostData(cr.RequestData.PostParams).
		Post()
	if cr.requestResponse.err != nil {
		// 	return err, resp.Raw
		panic(cr.requestResponse.err)
	}
	if cr.requestResponse.response.IsOk() {
		res = cr.requestResponse.response.Body
	}
	return res
}

//GetResponse the detail of request response
func (cr *CURL) GetResponse() *CurlResponse {
	return &cr.requestResponse
}