package export

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/zhaochangjiang/golang-utils/myutils"
)

//AccessZendao
type AccessZendao struct {
	CURL
	Userinfo     Userinfo
	SendParams   map[string]string
	Link         string
	Domain       string
	Params       *map[string]string
	RepsonseBody string
	Log          *Log
}

//ZenDAOResult
type ZenDAOResult struct {
}

//NewZenDAOResult
func NewZenDAOResult() *ZenDAOResult {
	return &ZenDAOResult{}
}

//NewAccessZendao
func NewAccessZendao() *AccessZendao {
	return &AccessZendao{}
}

//Run 请求zendao地址
func (zda *AccessZendao) Run() (zendaoResult *ZenDAOResult, err error) {

	zendaoResult = NewZenDAOResult()

	//组织参数
	err = zda.orgParams()
	if err != nil {
		return
	}

	zda.initBaseLink()

	curlRequestData := NewCurlRequestData()
	curlRequestData.PostParams = *zda.Params
	curlRequestData.URI = zda.URI
	curlRequestData.Headers = make(map[string]string)
	zda.callCurl(curlRequestData)

	if zda.RepsonseBody == "" {
		str := "业务方返回信息为空字符串"
		zda.Log.Write(LogError, "the bussiness return is null string!")
		err = errors.New(str)
		return
	}
	err = json.Unmarshal([]byte(zda.RepsonseBody), zendaoResult)
	return
}

//发送请求获取数据
func (zda *AccessZendao) callCurl(crd *CurlRequestData) {
	var err error
	var time = 3
	zda.CURL.URI = crd.URI
	zda.CURL.RequestData = crd
	if _, ok := zda.CURL.RequestData.Headers["Content-Type"]; ok {
	} else {
		zda.CURL.RequestData.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	}
	//未防止一次请求数据超时，发送3次请求
	for i := 0; i < time; i++ {
		zda.RepsonseBody, err = zda.CURL.Request()
		if nil == err { //如果有一次获取数据成功，那么就表示成功,退出循环和方法
			return
		}
		zda.Log.Write(LogError, "第 "+strconv.Itoa(i)+" 次尝试请求数据失败,内容描述"+err.Error())
	}
	zda.Log.Write(LogError, "连续尝试 "+strconv.Itoa(time)+" 次获取数据失败!")
}

func (zda *AccessZendao) orgParams() error {
	customUser := make(map[string]interface{})
	customUser["id"] = "--mock-user-id--"
	customUser["username"] = "mockuser"
	customUser["realname"] = "mockuser"
	customUser["roleid"] = ""
	organ := make(map[string]string)
	organ["orgroot"] = zda.Userinfo.Orgroot
	organ["orgcode"] = zda.Userinfo.Orgcode
	organ["customerId"] = zda.Userinfo.CustomerId
	organ["name"] = zda.Userinfo.Name
	organ["theme"] = zda.Userinfo.Theme
	customUser["organ"] = organ

	if customUserJSON, err := json.Marshal(customUser); err != nil {
		return err
	} else {
		(*zda.Params)["customuser"] = string(customUserJSON)
	}
	var sendParams = zda.SendParams

	if sendParamsJSON, err := json.Marshal(sendParams); err != nil {
		return err
	} else {
		(*zda.Params)["customparams"] = string(sendParamsJSON)
	}
	(*zda.Params)["format"] = "json"
	(*zda.Params)["timestamp"] = myutils.TimeNow()
	// $time                              = time();
	// $dateTimeString                    = date('Y-m-d H:i:s', $time);
	// $sendParams                        = $this->params['params']['remoteParams'];//传递的参数

	// $params                            = [
	// 	'customparams' => json_encode($sendParams),
	// 	'customuser'   => json_encode($customUser),
	// 	'method'       => "{$this->params['module']}.{$this->params['method']}",
	// 	'timestamp'    => $dateTimeString,
	// ];

	//签名生成
	zda.encryption()
	return nil
}

//GetEnvOldKey 获得Zendao老框架通信的key
func (zda *AccessZendao) GetEnvOldKey() string {
	return "cat"
}
func (zda *AccessZendao) ksort() {
	keyString := make([]string, 0)
	for k := range *(zda.Params) {
		keyString = append(keyString, k)
	}
	//字符串升序排序
	sort.Strings(keyString)
	param := make(map[string]string)
	for _, v := range keyString {
		param[v] = (*zda.Params)[v]
	}
	*zda.Params = param
}

//SetDomain
func (zda *AccessZendao) SetDomain(domain string) {
	zda.Domain = domain
}
func (zda *AccessZendao) initBaseLink() {
	zda.Link = zda.Domain + "/inside.php?t=json&m=index&f=service"
}

//encryption  签名算法
func (zda *AccessZendao) encryption() {
	//获得老框架的签名KEY
	appSec := zda.GetEnvOldKey()
	if appSec == "" {
		appSec = "cat"
	}
	sign := appSec
	zda.ksort()
	for k, v := range *zda.Params {
		if k != "" && v != "" {
			sign += k + v
		}
	}
	sign += appSec
	md5String := myutils.StringMd5EqualPHP(sign)
	(*zda.Params)["sign"] = strings.ToUpper(md5String)
}
