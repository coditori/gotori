package main

import (
	"log"
	"net/http"
	"os"
	"rest/middlewares"
	"rest/models"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.User).UserName,
		"text":     "Hello World.",
	})
}

// User demo

func main() {
	port := os.Getenv("PORT")
	r := gin.Default()

	if port == "" {
		port = "8000"
	}

	authMiddleware := middlewares.JwtAuth()

	// the jwt middleware

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth", authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		auth.GET("/hello", helloHandler)
	}
	// Refresh time can be longer than token timeout

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
