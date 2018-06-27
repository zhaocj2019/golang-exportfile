package export

import (
	"errors"
	"strconv"
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

//NewCurl 初始化curl请求
func NewCurl() *CURL {
	once.Do(func() {
		cr = &CURL{}
	})
	return cr
}

//Request 向服务器请求信息
func (cr *CURL) Request() (string, error) {

	reponseString := ""
	req := curl.NewRequest()
	cr.requestResponse.response, cr.requestResponse.err = req.
		SetUrl(cr.RequestData.URI).
		SetHeaders(cr.RequestData.Headers).
		SetCookies(cr.RequestData.Cookies).
		SetQueries(cr.RequestData.GetParams).
		SetPostData(cr.RequestData.PostParams).
		Post()
	if cr.requestResponse.err != nil {
		return reponseString, cr.requestResponse.err
	}
	if cr.requestResponse.response.IsOk() {
		reponseString = cr.requestResponse.response.Body
	} else {
		//如果服务器返回的状态不是200,则报错
		err := errors.New("the server return status is wrong(http code:" + strconv.Itoa(cr.requestResponse.response.Raw.StatusCode) + ")")
		return reponseString, err
	}
	return reponseString, nil
}

//GetResponse the detail of request response
func (cr *CURL) GetResponse() *CurlResponse {
	return &cr.requestResponse
}
