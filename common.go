package main

import (
	"errors"
	"strings"
)

//类型转换

//配置jpush应用的秘钥
const (
	appKey       = "########################"
	masterSecret = "########################"
)

//Json返回的信息
type ReturnMessage struct {
	Status int         `json:"status"`
	Info   interface{} `json:"info"`
	Url    string      `json:"url,omitempty"`
}

//业务返回码
var HttpStatusCode map[int]string = map[int]string{
	200: "Success!",
	400: "错误的请求!",
	401: "未验证!",
	403: "被拒绝!",
	404: "无法找到!",
	405: "请求方法不合适!",
	410: "已下线!",
	429: "过多的请求!",
	500: "内部服务器错误!",
	502: "无效的代理!",
	503: "服务暂时失效!",
	504: "代理超时!",
}

/*=============过滤函数=============*/

//过滤Ttile
func ReturnTitle(title string) string {
	return string(title)
}

//过滤Audience 格式如下0_all 1_shue,roy
func ReturnAudience(audience string) (err error, style string, result interface{}) {
	errMsg := "第二个参数格式错误！格式如下0_all 1_shue,roy"
	//分割字符串 0_string,string
	arrayStr := strings.Split(audience, "_")
	//保证分割的字符串有_
	if len(arrayStr) >= 2 && arrayStr[1] != "" {
		//判断arrayStr[]
		switch arrayStr[0] {
		case "0":
			if arrayStr[1] != "all" {
				err = errors.New(errMsg)
				result = nil
			} else {
				err = nil
				result = "all"
			}
		case "1", "2":
			//继续分割字符串
			if arrayStr[1] == "" {
				err = errors.New(errMsg)
				result = nil
			} else {
				arrayStrChild := strings.Split(arrayStr[1], ",")
				err = nil
				result = arrayStrChild
			}
		default:
			err = errors.New(errMsg)
			result = nil
		}
	} else {
		err = errors.New(errMsg)
		result = nil
	}
	style = arrayStr[0]
	return
}

//过滤Extras Extras格式 url_value
func ReturnExtras(extras string) (err error, result interface{}) {

	if extras == "" {
		return nil, make(map[string]interface{})
	}

	errMsg := "第三个参数格式错误！格式如下url,text_baidu.com,go"
	//分割字符串 url,text_baidu.com,go
	arrayStr := strings.Split(extras, "_")

	if len(arrayStr) >= 2 {
		//初次化map用于返回数据
		mapExtras := make(map[string]interface{})
		arrayStrChildKey := strings.Split(arrayStr[0], ",")
		arrayStrChildValue := strings.Split(arrayStr[1], ",")
		for key, value := range arrayStrChildKey {
			mapExtras[value] = arrayStrChildValue[key]
		}
		err = nil
		result = mapExtras

	} else {
		err = errors.New(errMsg)
		result = nil
	}
	return
}
