package main

import (
	"github.com/DeanThompson/jpush-api-go-client"
	"github.com/DeanThompson/jpush-api-go-client/push"
)

type Jpush struct {
	title    string                 //推送标题
	style    string                 //0是全部对象，1是别名，2是注册id
	audience interface{}            //推送对象 0_all
	extras   map[string]interface{} //推送附加字段用于逻辑业务判断
}

//对象初始化函数
func NewJpush(title string, style string, audience interface{}, extras map[string]interface{}) *Jpush {
	return &Jpush{title, style, audience, extras}
}

//推送消息
func (j *Jpush) PushDevice() (err error, result *push.PushResult) {
	// platform 对象
	platform := push.NewPlatform()
	// 用 Add() 方法添加具体平台参数，可选: "all", "ios", "android"
	platform.Add("ios")

	// audience 对象，表示消息受众
	audience := push.NewAudience()
	// 和 platform 一样，可以调用 All() 方法设置所有受众
	switch j.style {
	case "0":
		audience.All()
	case "1":
		audience.SetAlias(j.audience.([]string))
	case "2":
		audience.SetRegistrationId(j.audience.([]string))
	default:
		audience.All()
	}

	// notification 对象，表示 通知，传递 alert 属性初始化
	notification := push.NewNotification(j.title)

	// iOS 平台专有的 notification，用 alert 属性初始化
	iosNotification := push.NewIosNotification(j.title)
	iosNotification.Badge = 1

	//这里自定义 Key/value 信息，以供业务使用。
	if len(j.extras) > 0 {
		extrasMap := j.extras
		iosNotification.Extras = extrasMap
	}
	notification.Ios = iosNotification

	message := push.NewMessage(j.title)
	message.Title = j.title

	// option 对象，表示推送可选项
	options := push.NewOptions()
	// Options 的 Validate 方法会对 time_to_live 属性做范围限制，以满足 JPush 的规范
	options.TimeToLive = 86400
	// iOS 平台，是否推送生产环境，false 表示开发环境；如果不指定，就是生产环境
	options.ApnsProduction = true
	// 构建jpush推送的数据对象
	payload := push.NewPushObject()
	payload.Platform = platform
	payload.Audience = audience
	payload.Notification = notification
	payload.Message = message
	payload.Options = options

	//创建jpush对象
	jclient := jpush.NewJPushClient(appKey, masterSecret)

	//Push 会推送到客户端
	result, err = jclient.Push(payload)

	// PushValidate 的参数和 Push 完全一致
	// 区别在于，PushValidate 只会验证推送调用成功，不会向用户发送任何消息
	//result, err = jclient.PushValidate(payload)

	if err != nil {
		return err, result
	} else {
		return nil, result
	}
}
