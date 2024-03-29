package server

import (
	"log"
	"rest/middlewares"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"
var authMiddleware *jwt.GinJWTMiddleware = middlewares.JwtAuth()

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(middlewares.CORSMiddleware())
	// router.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	router.GET("/ping", pingHandler)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/login", loginHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := router.Group("/auth", authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.POST("/videos", videoController.Save)
		auth.GET("/videos", videoController.FindAll)
		auth.GET("/videos/:id", videoController.FindById)
		auth.PUT("/videos/:id", videoController.Update)
		auth.DELETE("/videos/:id", videoController.Delete)
	}

	return router
}

// GetPing     godoc
// @Summary      Ping the API
// @Description  Return poing if server is fine.
// @Tags         ping
// @Produce      json
// @Success      200    {object} object
// @Router       /ping [get]
func pingHandler(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"messsage": "pong!",
	})
}

// PostLogin     godoc
// @Summary      Login by username and password
// @Description  Takes nothing. Return list of videos.
// @Tags         login
// @Produce      json
// @Param        login  body      models.LoginRequest  true  "Login JSON"
// @Success      200    {object}  object
// @Router       /login [post]
func loginHandler(ctx *gin.Context) {
	authMiddleware.LoginHandler(ctx)
}
