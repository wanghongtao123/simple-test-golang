package main

import (
	"net/http"
	"time"
	"fmt"
	"strings"
	"github.com/gin-gonic/gin"
)

var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// 跨域
func Cors() gin.HandlerFunc {
    return func(c *gin.Context) {
        method := c.Request.Method      //请求方法
        origin := c.Request.Header.Get("Origin")        //请求头部
        var headerKeys []string                             // 声明请求头keys
        for k, _ := range c.Request.Header {
            headerKeys = append(headerKeys, k)
        }
        headerStr := strings.Join(headerKeys, ", ")
        if headerStr != "" {
            headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
        } else {
            headerStr = "access-control-allow-origin, access-control-allow-headers"
        }
        if origin != "" {
            c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
            c.Header("Access-Control-Allow-Origin", "*")        // 这是允许访问所有域
            c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")      //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
            //  header的类型
            c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
            //              允许跨域设置                                                                                                      可以返回其他子段
            c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")      // 跨域关键设置 让浏览器可以解析
            c.Header("Access-Control-Max-Age", "172800")        // 缓存请求信息 单位为秒
            c.Header("Access-Control-Allow-Credentials", "false")       //  跨域请求是否需要带cookie信息 默认设置为true
            c.Set("content-type", "application/json")       // 设置返回格式是json
        }

        //放行所有OPTIONS方法
        if method == "OPTIONS" {
            c.JSON(http.StatusOK, "Options Request!")
        }
        // 处理请求
        c.Next()        //  处理请求
    }
}



func main() {
	// 获得当前gin框架引擎，对外主要是router
	engine := gin.Default()
	engine.Use(Cors())
	//  完成各项功能的router注册

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	
	// 提供字面字符
	engine.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	authorized := engine.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})


	s := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
	// listen and serve on 0.0.0.0:8080
}
