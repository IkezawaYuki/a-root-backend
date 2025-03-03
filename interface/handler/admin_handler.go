package handler

import (
	"IkezawaYuki/a-root-backend/interface/dto/req"
	_ "IkezawaYuki/a-root-backend/interface/dto/res"
	"IkezawaYuki/a-root-backend/usecase"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	adminUsecase usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase usecase.AdminUsecase) AdminHandler {
	return AdminHandler{
		adminUsecase: adminUsecase,
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

// DeleteAdmin
// @Summary 管理者ユーザーを削除する
// @Description
// @Tags admin
// @Accept application/json
// @Produce application/json
// @Success 202 {object} res.Message
// @Router /admin/{admin_id} [DELETE]
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
// @Router /admin/{admin_id} [PUT]
func (h AdminHandler) UpdateAdmin(c *gin.Context) {
	slog.Info("UpdateAdmin is invoked")
	var body req.UpdateAdminBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	p := c.Param("admin_id")
	adminID, err := strconv.Atoi(p)
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
	p := c.Param("customer_id")
	customerID, err := strconv.Atoi(p)
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
