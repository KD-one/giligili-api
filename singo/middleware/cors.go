package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
			//
			//c.Header("Content-Type", "application/json")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.AbortWithStatus(http.StatusNoContent)
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
	//config := cors.DefaultConfig()
	//config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	//config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	//if gin.Mode() == gin.ReleaseMode {
	//	// 生产环境需要配置跨域域名，否则403
	//	config.AllowOrigins = []string{"http://121.199.8.131:80", "http://localhost:8080"}
	//} else {
	//	// 测试环境下模糊匹配本地开头的请求
	//	config.AllowOriginFunc = func(origin string) bool {
	//		if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
	//			return true
	//		}
	//		if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
	//			return true
	//		}
	//		return false
	//	}
	//}
	//config.AllowCredentials = true
	//return cors.New(config)
}
