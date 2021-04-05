package main

import (
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

	accessLogicImpl := urlMetric.AccessLogicImpl{}

	route := gin.Default()
	route.Use(Cors())

	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	route.GET("/pageMetric", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		accessMetric := accessLogicImpl.GetUrlMetric(url)
		c.JSON(http.StatusOK, &accessMetric)
	})

	route.GET("/browse").GET("/page", func(c *gin.Context) {
		url := c.Query("url")
		if url == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		accessMetric := urlMetric.AccessMetric{}
		accessMetric.Url = url
		accessMetric.BrowseUserCnt = 1
		accessLogicImpl.PutUrlMetric(&accessMetric)

		c.Status(http.StatusOK)
	})

	e := route.Run()

	if e != nil {
		panic(e)
	}
}
