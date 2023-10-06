package api

import (
	_ "playground/cpp-bootcamp/api/docs"
	"playground/cpp-bootcamp/api/handler"
	"playground/cpp-bootcamp/pkg/helper"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewServer(h *handler.Handler) *gin.Engine {
	r := gin.Default()
	r.Use(helper.Logger)

	r.POST("/login", h.Login,helper.Logger)
	r.POST("/user", h.CreateUser)
	r.Use(helper.EndMiddleware)
	r.GET("/user", h.GetAllUsers)
	r.GET("/user/:id", h.GetUser)
	r.PUT("/user/:id", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
