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
		c.JSON(http.StatusNotFound, err.Error())
	case errors.Is(err, arootErr.ErrAuthentication):
		c.JSON(http.StatusForbidden, err.Error())
	case errors.Is(err, arootErr.ErrAuthorization):
		c.JSON(http.StatusUnauthorized, err.Error())
	case errors.Is(err, arootErr.ErrDuplicateKey):
		c.JSON(http.StatusConflict, err.Error())
	case errors.Is(err, arootErr.ErrDuplicateEmail):
		c.JSON(http.StatusConflict, err.Error())
	case errors.Is(err, arootErr.ErrEmailUsed):
		c.JSON(http.StatusBadRequest, err.Error())
	default:
		c.JSON(http.StatusInternalServerError, err.Error())
	}
}
