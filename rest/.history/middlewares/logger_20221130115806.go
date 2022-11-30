package middlewares

import "github.com/gin-gonic/gin"

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func() {param gin.LogFormatterParams})
}
