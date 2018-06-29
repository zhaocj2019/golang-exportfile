package export

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/zhaochangjiang/golang-utils/myutils"
)

//NewAccessAlpha 创建实例
func NewAccessAlpha() *AccessAlpha {
	return &AccessAlpha{Type: "VEGA", Method: "POST", ConnectTimeout: 3, ResponseTimeout: 30}
}

//AccessAlpha 访问Alpha框架
type AccessAlpha struct {
	CURL
	RepsonseBody    string
	Log             *Log
	Type            string //Enum("vega","zuul")
	ConnectTimeout  int
	ResponseTimeout int
	Method          string //Enum("POST","GET")
	Params          *map[string]string
	API             string
}

func (e *AccessAlpha) runZuul() (*AlphaResult, error) {
	var err error
	e.encryption()
	curlRequestData := NewCurlRequestData()
	curlRequestData.PostParams = *e.Params
	curlRequestData.URI = e.URI
	curlRequestData.Headers = make(map[string]string)

	e.callCurl(curlRequestData)

	var alphaResult = NewAlphaResult()
	if e.RepsonseBody != "" {
		err = json.Unmarshal([]byte(e.RepsonseBody), alphaResult)
	} else {
		str := "业务方返回信息为空字符串"
		e.Log.Write(LogError, "the bussiness return is null string!")
		err = errors.New(str)
		return nil, err
	}
	if alphaResult.Code > 0 {
		jsonString, err := json.Marshal(curlRequestData)
		if err != nil {
			e.Log.Write(LogError, "the format of the return is not a json string!")
			return nil, err
		}
		e.Log.Write(LogError, " ERROR MESSAGE:"+alphaResult.Msg+" PARAMS:"+string(jsonString)+" RETURN:"+e.RepsonseBody)
	}
	return alphaResult, err
}

//Run 返回结果
func (e *AccessAlpha) Run() (alphaResult *AlphaResult, err error) {

	e.initLinkAddress()
	switch e.Type {
	case "ZUUL":
		alphaResult, err = e.runZuul()
		break
	case "VEGA":
		alphaResult, err = e.runVega()
		break
	default:
		panic("the params type is error!")
	}
	return
}
func (e *AccessAlpha) runVega() (alphaResult *AlphaResult, err error) {

	alphaResult = NewAlphaResult()

	return
}

//AlphaResult alpha框架返回的内容
type AlphaResult struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	SubCode int         `json:"sub_code"`
	SubMsg  string      `json:"sub_msg"`
	ReqID   string      `json:"req_id"`
	Data    interface{} `json:"data"`
}

//NewAlphaResult
func NewAlphaResult() *AlphaResult {
	return &AlphaResult{}
}

//发送请求获取数据
func (e *AccessAlpha) callCurl(crd *CurlRequestData) {
	var err error
	var time = 3
	e.CURL.URI = crd.URI
	e.CURL.RequestData = crd
	if _, ok := e.CURL.RequestData.Headers["Content-Type"]; ok {
	} else {
		e.CURL.RequestData.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	//未防止一次请求数据超时，发送3次请求
	for i := 0; i < time; i++ {
		e.RepsonseBody, err = e.CURL.Request()
		if nil == err { //如果有一次获取数据成功，那么就表示成功,退出循环和方法
			return
		}
		e.Log.Write(LogError, "第 "+strconv.Itoa(i)+" 次尝试请求数据失败,内容描述"+err.Error())
	}
	e.Log.Write(LogError, "连续尝试 "+strconv.Itoa(time)+" 次获取数据失败!")
}

//SetParams 设置请求参数
func (e *AccessAlpha) SetParams(params *map[string]string) *AccessAlpha {
	e.Params = params
	return e
}

//拼接应用路径
func (e *AccessAlpha) pathSpell(api string) (string, error) {
	var resString = ""
	pathArray := strings.Split(api, ".")
	splitString := "-"
	version := ""
	product := ""
	divString := "/"

	if len(pathArray) != 3 {
		err := errors.New("你发送的请求地址:" + api + " 不正确")
		return resString, err
	}

	other := strings.Split(pathArray[0], splitString)
	productLength := len(other)
	if productLength < 2 { //此处只判断小于2报错
		err := errors.New("你发送的请求地址:" + api + " 不正确")
		return resString, err
	}

	for k, v := range other {
		if k == 0 {
			product = v
		} else if k == productLength-1 { //最后一条为version版本信息
			version = v
		} else {
			product += splitString + v
		}
	}
	if strings.Index(version, "v") != 0 {
		err := errors.New("你发送的请求地址:" + api + " 版本号不正确")
		return resString, err
	}

	//拼接应用内容
	returnString := divString + version + divString + product + divString + pathArray[1] + divString + pathArray[2]
	return returnString, nil
}

//SetAPI 设置请求参数
func (e *AccessAlpha) SetAPI(api string) *AccessAlpha {
	var err error
	e.API, err = e.pathSpell(api)
	if nil != err {
		panic(err)
	}
	return e
}

//SetType 设置请求类型
func (e *AccessAlpha) SetType(typeString string) *AccessAlpha {
	e.Type = strings.ToUpper(typeString)
	switch e.Type {
	case "VEGA":
	case "ZUUL":
		break
	default:
		panic("当前只支持vega 和zuul访问")
	}
	return e
}

//SetMethod 设置访问请求的方式
func (e *AccessAlpha) SetMethod(method string) *AccessAlpha {
	e.Method = strings.ToUpper(method)
	switch e.Method {
	case "POST":
	case "GET":
		break
	default:
		panic("当前只支持POST 和GET访问")
	}
	return e
}

type accessConfig struct {
	Host      string
	Port      int
	accessid  string
	secretkey string
}

//GetAccessConfig 或得访问网关的配置
func (e *AccessAlpha) GetAccessConfig() *accessConfig {
	access := &accessConfig{accessid: "buggf8c", secretkey: "raUAPV6cpx17jKcsOMQmPazvf5KwWTfI", Host: "http://172.22.34.203", Port: 5555}
	// access = &accessConfig{accessid: "buggf8c", secretkey: "Wgyokg4EC2dJKwOYwBu3zrZHgfnOYHWi", Host: "http://172.22.34.203", Port: 5555}
	// access = &accessConfig{accessid: "buggf8c", secretkey: "raUAPV6cpx17jKcsOMQmPazvf5KwWTfI", Host: "http://test.org", Port: 80}
	//http://172.22.34.203:5555/v1/path_planning/truck/index?accessid=b167byj&sign=fSuKs38XvFEHH6iSwdmpuFKcxEg%3D&token=c071d9d753ab0c26e62e0b80d789d09d&g7timestamp=1505195584540&pageNo=1&pageSize=10&start=2017-09-05+00:00:00&end=2017-09-12+23:59:59&customer=&average=0
	return access
}
func (e *AccessAlpha) initLinkAddress() {
	accessConfig := e.GetAccessConfig()

	//默认访问地址
	if accessConfig.Host == "" {
		accessConfig.Host = "http://vega.huoyunren.com"
	}

	//默认访问端口
	if accessConfig.Port == 0 {
		accessConfig.Port = 80
	}
	e.URI = accessConfig.Host + ":" + strconv.Itoa(accessConfig.Port) + e.API
}

//实例 http://172.22.34.203:5555/v1/path_planning/truck/index?accessid=b167byj&sign=fSuKs38XvFEHH6iSwdmpuFKcxEg%3D&token=c071d9d753ab0c26e62e0b80d789d09d&g7timestamp=1505195584540&pageNo=1&pageSize=10&start=2017-09-05+00:00:00&end=2017-09-12+23:59:59&customer=&average=0
//签名生成方发
func (e *AccessAlpha) encryption() {

	//获得单位为毫秒的UNIX时间戳
	unixTimstampString := strconv.FormatInt(int64(time.Now().UnixNano()/1000000), 10)
	// unixTimstampString = "1505195584540"
	(*e.Params)["g7timestamp"] = unixTimstampString
	StringToSign := strings.ToUpper(e.Method) + "\n" + unixTimstampString + "\n" + e.API
	// fmt.Println("StringToSign=", StringToSign)
	// fmt.Println("CanonicalizedResource=", e.API)
	config := e.GetAccessConfig()
	Signature := myutils.StringBase64Encode(myutils.StringHMACSHA1(myutils.StringUTF8EncodingOf([]byte(StringToSign), "UTF-8"), config.secretkey))
	(*e.Params)["sign"] = string(Signature)
	// fmt.Println("Sign=", (*e.Params)["sign"] )
	(*e.Params)["accessid"] = config.accessid

}
