package handler

import (
	"IkezawaYuki/a-root-backend/domain/entity"
	"IkezawaYuki/a-root-backend/interface/dto/req"
	_ "IkezawaYuki/a-root-backend/interface/dto/res"
	"IkezawaYuki/a-root-backend/interface/session"
	"github.com/redis/go-redis/v9"

	//"IkezawaYuki/a-root-backend/interface/middleware"
	"IkezawaYuki/a-root-backend/usecase"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
	redisClient  *redis.Client
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase, redisClient *redis.Client) AdminHandler {
	return AdminHandler{
		adminUsecase: adminUsecase,
		redisClient:  redisClient,
	}
}

// CreateAdmin
// @Summary 管理者ユーザーを作成する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 201 {object} res.Admin
// @Router /admin/admins [POST]
func (h AdminHandler) CreateAdmin(c *gin.Context) {
	slog.Info("CreateAdmin is invoked")
	var body req.CreateAdminBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.CreateAdmin(c.Request.Context(), body)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, resp)
}

// GetAdmins
// @Summary 管理者ユーザー一覧を取得する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param email query string false "メールアドレス"
// @param partialName query string false "名前（部分一致）"
// @param limit query int false "取得件数"
// @param offset query int false "取得位置"
// @Success 200 {object} res.Admins
// @Router /admin/admins [GET]
func (h AdminHandler) GetAdmins(c *gin.Context) {
	slog.Info("GetAdmins is invoked")
	var query req.AdminQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.GetAdmins(c.Request.Context(), query)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetAdmin
// @Summary 管理者ユーザーを取得する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param admin_id path string true "管理者ID"
// @Success 200 {object} res.Admins
// @Router /admin/admins/{admin_id} [GET]
func (h AdminHandler) GetAdmin(c *gin.Context) {
	slog.Info("GetAdmin is invoked")
	adminID, err := strconv.Atoi(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.GetAdmin(c.Request.Context(), adminID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteAdmin
// @Summary 管理者ユーザーを削除する
// @Description
// @Tags admin
// @param admin_id path string true "管理者ID"
// @Accept application/json
// @Produce application/json
// @Success 202 {object} res.Message
// @Router /admin/admins/{admin_id} [DELETE]
func (h AdminHandler) DeleteAdmin(c *gin.Context) {
	slog.Info("DeleteAdmin is invoked")
	p := c.Param("admin_id")
	adminID, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.DeleteAdmin(c.Request.Context(), adminID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// UpdateAdmin
// @Summary 管理者ユーザーを更新する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 202 {object} res.Admin
// @Router /admin/admins/{admin_id} [PUT]
func (h AdminHandler) UpdateAdmin(c *gin.Context) {
	slog.Info("UpdateAdmin is invoked")
	var body req.UpdateAdminBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	adminID, err := strconv.Atoi(c.Param("admin_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.UpdateAdmin(c.Request.Context(), adminID, body)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusAccepted, resp)
}

// GetCustomers
// @Summary 顧客一覧を取得する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param email query string false "メールアドレス"
// @param partialName query string false "名前（部分一致）"
// @param IsFacebookToken query bool false "フェイスブック連携"
// @param limit query int false "取得件数"
// @param offset query int false "取得位置"
// @Success 200 {object} res.Customers
// @Router /admin/customers [GET]
func (h AdminHandler) GetCustomers(c *gin.Context) {
	slog.Info("GetCustomers is invoked")
	var query req.CustomerQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.GetCustomers(c.Request.Context(), query)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// GetCustomer
// @Summary 顧客情報を取得する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Customer
// @Router /admin/customers/{customer_id} [GET]
func (h AdminHandler) GetCustomer(c *gin.Context) {
	slog.Info("GetCustomer is invoked")
	customerID, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.GetCustomer(c.Request.Context(), customerID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateCustomer
// @Summary 顧客情報を更新する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Customer
// @Router /admin/customers/{customer_id} [PUT]
func (h AdminHandler) UpdateCustomer(c *gin.Context) {
	slog.Info("UpdateCustomer is invoked")
	customerID, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var body req.UpdateCustomerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.UpdateCustomer(c.Request.Context(), customerID, body)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// CreateCustomer
// @Summary 顧客情報を登録する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Customer
// @Router /admin/customers [POST]
func (h AdminHandler) CreateCustomer(c *gin.Context) {
	slog.Info("CreateCustomer is invoked")

	var body req.CreateCustomerBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.CreateCustomer(c.Request.Context(), body)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// DeleteCustomer
// @Summary 顧客情報を削除する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 200 {object} res.Customer
// @Router /admin/customers/{customer_id} [DELETE]
func (h AdminHandler) DeleteCustomer(c *gin.Context) {
	slog.Info("DeleteCustomer is invoked")
	customerID, err := strconv.Atoi(c.Param("customer_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.DeleteCustomer(c.Request.Context(), customerID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

// Login
// @Summary ログインする
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @param default body req.User true "ログイン情報"
// @Success 200 {object} res.Auth
// @Router /customer/login [POST]
func (h AdminHandler) Login(c *gin.Context) {
	slog.Info("Login is invoked")
	var user req.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.adminUsecase.Login(c.Request.Context(), user)
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
