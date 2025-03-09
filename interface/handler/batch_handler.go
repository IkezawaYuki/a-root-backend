package handler

import (
	"IkezawaYuki/a-root-backend/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BatchHandler struct {
	batchUsecase usecase.BatchUsecase
}

func NewBatchHandler(batchUsecase usecase.BatchUsecase) BatchHandler {
	return BatchHandler{
		batchUsecase: batchUsecase,
	}
}

// SyncInstagramToWordPress
// @Summary [Instagram=>WordPress]連携を実行する
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Message
// @Router /batch/sync [POST]
func (h BatchHandler) SyncInstagramToWordPress(c *gin.Context) {
	resp, err := h.batchUsecase.SyncInstagramToWordPress(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// RefreshToken
// @Summary トークンを更新する
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Message
// @Router /batch/refresh [POST]
func (h BatchHandler) RefreshToken(c *gin.Context) {
	resp, err := h.batchUsecase.RefreshToken(c.Request.Context())
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
