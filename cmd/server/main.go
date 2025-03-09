package main

import (
	"IkezawaYuki/a-root-backend/di"
	_ "IkezawaYuki/a-root-backend/docs"
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title a-root-backend
// @version 1.0.0
// @description a-root-backend
// @contact.url https://github.com/IkezawaYuki/a-root-backend
// @BasePath /v1
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5137"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	di.Connection()
	customerHandler := di.NewCustomerHandler()
	adminHandler := di.NewAdminHandler()

	v1 := r.Group("/v1")
	{
		customer := v1.Group("/customer")
		{
			customer.GET("/me", customerHandler.GetMe)
		}
		admin := v1.Group("/admin")
		{
			admin.GET("/admins", adminHandler.GetAdmins)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	fmt.Println("----------")
	fmt.Println("a-root-backend server started")
	fmt.Println("----------")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	di.Close()
	log.Println("Server exiting")
}
