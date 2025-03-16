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
	customerID := c.MustGet("customer_id").(int)
	resp, err := h.customerUsecase.GetCustomer(c.Request.Context(), customerID)
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
	customerID := c.MustGet("customer_id").(int)
	var query req.PostQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.customerUsecase.GetPosts(c.Request.Context(), customerID, query)
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
	customerID := c.MustGet("customer_id").(int)
	resp, err := h.customerUsecase.FetchInstagramPosts(c.Request.Context(), customerID)
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
	customerID := c.MustGet("customer_id").(int)
	resp, err := h.customerUsecase.FetchAndPost(c.Request.Context(), customerID)
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
	err = session.SetLoginSession(c, entity.ARootCustomer, h.redisClient, resp.UserID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// TempRegister
// @Summary ユーザーを登録する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param default body req.EmailBody true "メールアドレス"
// @Success 201 {object} res.Message
// @Router /customer/temp_register [POST]
func (h CustomerHandler) TempRegister(c *gin.Context) {
	slog.Info("TempRegister is invoked")
	var email req.EmailBody
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.customerUsecase.TempRegister(c.Request.Context(), email)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// CheckToken
// @Summary トークンの検証
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param default body req.Token true "トークン"
// @Success 201 {object} res.Customer
// @Router /customer/check_token [POST]
func (h CustomerHandler) CheckToken(c *gin.Context) {
	slog.Info("CheckToken is invoked")
	var token req.Token
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.customerUsecase.CheckToken(c.Request.Context(), token.Token)
	if err != nil {
		handleError(c, err)
		return
	}
	err = session.SetLoginSession(c, entity.ARootCustomer, h.redisClient, resp.ID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// Register
// @Summary ユーザーを登録する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param default body req.RegisterCustomer true "メールアドレス"
// @Success 201 {object} res.Message
// @Router /customer/register/ [POST]
func (h CustomerHandler) Register(c *gin.Context) {
	slog.Info("Register is invoked")
	customerID := c.MustGet("customer_id").(int)
	var customer req.RegisterCustomer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.customerUsecase.Register(c.Request.Context(), customerID, customer)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}
