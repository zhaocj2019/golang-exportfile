package data

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	utils "github.com/zhaochangjiang/golang-utils/utils"
)

//ExportAbstract 公共参数解析
type ExportAbstract struct {
	httpRequestContent *http.Request
	RequestParams      interface{}
}

//错误信息收集
func (ea *ExportAbstract) Error() {
	if err := recover(); err != nil {
		fmt.Println(err) // 这里的err其实就是panic传入的内容，55
	}
}

//New 初始化导出服务
func (ea *ExportAbstract) New(r *http.Request) *ExportAbstract {
	ea.initParams(r)
	defer ea.Error()
	return ea
}

//InitParams
func (ea *ExportAbstract) initParams(r *http.Request) {
	ea.httpRequestContent = r
	ea.httpRequestContent.ParseForm()
	ea.paramsOrganization()
}

func (ea *ExportAbstract) paramsMaps(k string, v []string, params *map[string]interface{}) {

	regex := regexp.MustCompile(`(\[.*\])+$`).FindAllString(k, -1)
	if len(regex) > 0 {
		switch len(regex) {
		case 1:
			count := len([]rune(regex[0]))
			key := utils.SubString(k, -count, count)
			regex[0] = strings.TrimRight(strings.TrimLeft(regex[0], "["), "]")
			list := strings.Split(regex[0], "][")
			len := len(list)
			if len > 1 {
				var res interface{}
				for m := len - 1; m < 0; m-- {
					res = ea.orgDataFormat(m, list, v, res)
				}
				params[key] = res
			}
			break
		default:
			panic("the params is not support,please do it. the content is follow:")
		}
		return params
	

}
func (ea *ExportAbstract) orgDataFormat(m int, list []string, v string, res interface{}) interface{} {
	keyVal, err := strconv.Atoi(list[m])
	if nil != err {
		panic(err)
	}
	var tmp interface{}
	if list[m] != "0" && keyVal == 0 {
		tmp = map[string]string{list[m]: v}
	} else {
		tmp = []string{v}
	}
	return tmp
}

//ParamsOrganization
func (ea *ExportAbstract) paramsOrganization() *map[string]interface{} {

	var params map[string]interface{}
	var c = ea.httpRequestContent
	if nil != c.Form {
		for k, v := range c.Form {
			ea.paramsMaps(k, v, &params)
			params[k] = v
		}
	}
	if nil != c.PostForm {
		for k, v := range c.PostForm {
			params[k] = v
		}
	}
	return &params
}
