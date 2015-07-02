package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

func main() {

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	//写入日志
	os.Mkdir("push_log", os.ModeDir)
	os.Chdir("push_log")
	/*go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		timeLog := time.NewTicker(60 * time.Second)
		for {
			select {
			case <-timeLog.C:
				yourfile, _ := os.Create("log_" + strconv.Itoa(r.Intn(100)) + ".log")
				router.Use(gin.LoggerWithWriter(yourfile))
			}
		}
	}()*/
	yourfile, _ := os.Create("server" + time.Now().Format("2006-01-02") + ".log")
	router.Use(gin.LoggerWithWriter(yourfile))

	//默认生成jpush的推送对象
	//jpush := NewJpush("量子股或有100个涨停", 0, "all", make(map[string]interface{}))

	//默认路由
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "欢迎来到量子股推送")
	})

	//推送路由群组
	push := router.Group("/push")
	{
		//推送设备
		push.GET("/device", func(c *gin.Context) {

			//推送的标题 title
			requestTitle := c.Query("title")
			if requestTitle == "" {
				c.JSON(http.StatusOK, ReturnMessage{0, "推送标题title不能为空", "/"})
				return
			}
			title := ReturnTitle(requestTitle)

			//推送对象 audience  格式是num_value  0_all 1_shuye,ben

			requestAudience := c.Query("audience")
			if requestAudience == "" {
				c.JSON(http.StatusOK, ReturnMessage{0, "推送对象audience不能为空", "/"})
				return
			}
			err, sTyle, result := ReturnAudience(requestAudience)
			if err != nil {
				c.JSON(http.StatusOK, ReturnMessage{0, err.Error(), "/"})
				return
			}
			/*=====推送对象的类型 0是全部对象，1是别名，2是注册id=====*/
			style := sTyle
			audience := result

			//推送额外字段 extras url_value url,text_www.baidu.com,go
			requestExtras := c.Query("extras")

			err, result = ReturnExtras(requestExtras)
			if err != nil {
				c.JSON(http.StatusOK, ReturnMessage{0, err.Error(), "/"})
				return
			}
			extras := result.(map[string]interface{})

			/*=============发送推送===========*/

			jpush := Jpush{title, style, audience, extras}
			err, info := jpush.PushDevice()

			if err != nil {
				//发生未知错误，如网络连接失败等，
				c.JSON(http.StatusOK, ReturnMessage{0, err, "/"})
			} else {
				if info.StatusCode == 200 {
					//推送成功
					c.JSON(http.StatusOK, ReturnMessage{1, HttpStatusCode[info.StatusCode], "/"})
				} else {
					//推送失败
					c.JSON(http.StatusOK, ReturnMessage{0, HttpStatusCode[info.StatusCode], "/"})
				}
			}
		})
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, ReturnMessage{0, "请访问正确的路由", "/"})
	})

	/*8080端口*/
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
