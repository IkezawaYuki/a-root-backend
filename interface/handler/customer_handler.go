package handler

import (
	"IkezawaYuki/a-root-backend/domain/entity"
	"IkezawaYuki/a-root-backend/interface/dto/req"
	_ "IkezawaYuki/a-root-backend/interface/dto/res"
	"IkezawaYuki/a-root-backend/interface/session"
	"IkezawaYuki/a-root-backend/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"net/http"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
	redisClient     *redis.Client
}

func NewCustomerHandler(customerUsecase usecase.CustomerUsecase, redisClient *redis.Client) CustomerHandler {
	return CustomerHandler{
		customerUsecase: customerUsecase,
		redisClient:     redisClient,
	}
}

// GetMe
// @Summary 自分の情報を取得する
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Customer
// @Router /customer/me [GET]
func (h CustomerHandler) GetMe(c *gin.Context) {
	slog.Info("GetMe is invoked")
	customerID, ok := c.Get(entity.UserSession)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid customer_id"})
		return
	}
	resp, err := h.customerUsecase.GetCustomer(c.Request.Context(), customerID.(int))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetPosts
// @Summary 連携済みの投稿データを取得する
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @param limit query int false "取得件数"
// @param offset query int false "取得位置"
// @Success 200 {object} res.Customer
// @Router /customer/posts [GET]
func (h CustomerHandler) GetPosts(c *gin.Context) {
	slog.Info("GetPosts is invoked")
	customerID, ok := c.Get("customer_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid customer_id"})
		return
	}
	var query req.PostQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.customerUsecase.GetPosts(c.Request.Context(), customerID.(int), query)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// FetchInstagramPosts
// @Summary インスタグラム上の投稿データを取得する
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Customer
// @Router /customer/instagram [GET]
func (h CustomerHandler) FetchInstagramPosts(c *gin.Context) {
	slog.Info("FetchInstagramPosts is invoked")
	customerID, ok := c.Get("customer_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid customer_id"})
		return
	}
	resp, err := h.customerUsecase.FetchInstagramPosts(c.Request.Context(), customerID.(int))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Sync
// @Summary インスタグラム上の投稿データをWordpressに連携する
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Message
// @Router /customer/sync [POST]
func (h CustomerHandler) Sync(c *gin.Context) {
	slog.Info("Sync is invoked")
	customerID, ok := c.Get("customer_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid customer_id"})
		return
	}
	resp, err := h.customerUsecase.FetchAndPost(c.Request.Context(), customerID.(int))
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Login
// @Summary ログインする
// @Description
// @Tags customer
// @Accept application/json
// @Produce application/json
// @param default body req.User true "ログイン情報"
// @Success 200 {object} res.Auth
// @Router /customer/login [POST]
func (h CustomerHandler) Login(c *gin.Context) {
	slog.Info("Login is invoked")
	var user req.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.customerUsecase.Login(c.Request.Context(), user)
	if err != nil {
		handleError(c, err)
		return
	}
	err = session.SetLoginSession(c, entity.ARootAdmin, h.redisClient, resp.UserID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}
