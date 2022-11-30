package middlewares

import "github.com/gin-gonic/gin"

func BasicAuth() gin.HandlerFunc {
	return gin.Accounts{
		"paragmatic": "reviews",
	}
}
