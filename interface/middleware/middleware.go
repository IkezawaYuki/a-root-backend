package middleware

import (
	"IkezawaYuki/a-root-backend/di"
	"IkezawaYuki/a-root-backend/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/rbcervilla/redisstore/v9"
	"net/http"
)

func validateSession(c *gin.Context, sessionName string) {
	client := di.NewRedisRepository().GetClient()
	store, err := redisstore.NewRedisStore(c.Request.Context(), client)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	sess, err := store.Get(c.Request, sessionName)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if sess.Values["uid"] == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "session expired",
		})
		return
	}
	uid, ok := sess.Values["uid"].(int)
	if !ok || uid == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "session expired",
		})
		return
	}
	c.Set(entity.UserSession, uid)
	c.Next()
}

func Customer(c *gin.Context) {
	validateSession(c, entity.ARootCustomer)
}

func Admin(c *gin.Context) {
	validateSession(c, entity.ARootAdmin)
}

func Batch(c *gin.Context) {
	c.Next()
}
