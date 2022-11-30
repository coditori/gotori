package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC822),
			param.Method,
			param.StatusCode,
			param.Latency,
		)
	})
}
