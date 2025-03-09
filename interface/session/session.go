package session

import (
	"github.com/gin-gonic/gin"
	"github.com/rbcervilla/redisstore/v9"
	"github.com/redis/go-redis/v9"
)

func SetLoginSession(c *gin.Context, sessionName string, client *redis.Client, userID int) error {
	store, err := redisstore.NewRedisStore(c.Request.Context(), client)
	if err != nil {
		return err
	}
	sess, err := store.Get(c.Request, sessionName)
	if err != nil {
		return err
	}
	sess.Values["uid"] = userID
	err = sess.Save(c.Request, c.Writer)
	if err != nil {
		return err
	}
	return nil
}
