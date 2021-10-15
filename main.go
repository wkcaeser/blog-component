package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
import urlMetric "blog-component/accessMetric"
import conf "blog-component/config"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if conf.GlobalConf.Cors.AccessControlAllowOrigin != "" {
			corsConf := conf.GlobalConf.Cors
			// 可将将* 替换为指定的域名
			c.Header("Access-Control-Allow-Origin", corsConf.AccessControlAllowOrigin)
			c.Header("Access-Control-Allow-Methods", corsConf.AccessControlAllowMethods)
			c.Header("Access-Control-Allow-Headers", corsConf.AccessControlAllowHeaders)
			c.Header("Access-Control-Expose-Headers", corsConf.AccessControlExposeHeaders)
			c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(corsConf.AccessControlAllowCredentials))
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}

func main() {

	accessLogImpl := urlMetric.AccessLogImpl{}

	route := gin.Default()
	route.Use(Cors())

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	route.GET("/log", func(c *gin.Context) {
		biz := c.Query("biz")
		if biz != "pv" {
			c.JSON(http.StatusNotFound, "biz invalid")
		}
		log := urlMetric.AccessLog{}

		log.ReqUri = c.Query("uri")
		log.UserUuid = c.Query("uuid")

		log.Ip = c.ClientIP()
		if log.Ip == "::1" {
			log.Ip = "127.0.0.1"
		}

		log.Site = c.GetHeader("origin")
		log.Uri = log.ReqUri[len(log.Site) : len(log.ReqUri)-1]

		marshal, _ := json.Marshal(c)
		log.RequestLog = string(marshal)

		accessLogImpl.PutUrlMetric(&log)
		pvMetric := accessLogImpl.GetUriMetric(log.Uri, log.UserUuid)
		c.JSON(http.StatusOK, &pvMetric)
	})

	e := route.Run("127.0.0.1:8080")

	if e != nil {
		panic(e)
	}
}
