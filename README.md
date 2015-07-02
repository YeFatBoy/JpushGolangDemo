# JpushGolangDemo
JpushGolangDemo


#Jpush And Golang
使用条件：首先得现在gin和jpush的封装包
https://github.com/gin-gonic/gin
https://github.com/DeanThompson/jpush-api-go-client

#Common.go
这里设置jpush平台的分配的key和密码还有一些公用的处理函数

#mian.go
gin的逻辑和路由等配置

#push.go
jpush的主程序

#device.go
更新设备信息,暂时没有开发

#summary
使用方法，运行之后再输入地址http://localhost:8080/push/device?title=测试&audience=1_SuperShuYe&extras=type_stock

title推送的标题必填
audience 推送对象，分几种模式

0_all推送所有的设备
1_shuye,shutou 推送多个设备别名
2_8de5gd,456de4d 推送多个注册id

extras 推送额外字段
url_http://www.baidu.com 单个
url,type_http://www.baidu.com,index 多个参数

