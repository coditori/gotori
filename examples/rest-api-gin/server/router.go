package server

import (
	"log"
	"rest/middlewares"
	"rest/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func NewRouter() *gin.Engine {
	router := gin.New()

	// router.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"messsage": "pong!",
		})
	})

	authMiddleware := middlewares.JwtAuth()

	router.POST("/login", authMiddleware.LoginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth", authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/hello", helloHandler)
		auth.GET("/videos", videoController.FindAll)
		auth.PUT("/videos/:id", videoController.Update)
		auth.DELETE("/videos/:id", videoController.Delete)
		auth.POST("/videos", videoController.Save)
	}

	return router
}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.User).UserName,
		"text":     "Hello World.",
	})
}
