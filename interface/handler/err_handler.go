package handler

import (
	"IkezawaYuki/a-root-backend/domain/arootErr"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, arootErr.ErrNotFound):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case errors.Is(err, arootErr.ErrAuthentication):
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	case errors.Is(err, arootErr.ErrAuthorization):
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	case errors.Is(err, arootErr.ErrDuplicateKey):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, arootErr.ErrDuplicateEmail):
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	case errors.Is(err, arootErr.ErrEmailUsed):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
