package middleware

import (
	"github.com/gin-gonic/gin"
)

func Customer(c *gin.Context) {
	c.Next()
}

func Admin(c *gin.Context) {
	c.Next()
}

func Batch(c *gin.Context) {
	c.Next()
}
