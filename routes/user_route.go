package routes

import (
	"api-test/handler"
	"api-test/helpers"
	"api-test/repositories"
	"api-test/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(group *gin.RouterGroup, db *gorm.DB) {
	repo := repositories.NewUserRepository(db)
	svc := services.NewUserService(repo)
	handler := handler.NewUserHandler(svc)

	group.POST("/register", handler.Register)
	group.POST("/login", handler.Login)

	auth := group.Use(helpers.JWTMiddleware())
	auth.GET("/:id", handler.GetProfile)
	auth.GET("", handler.GetAllUser)
	auth.DELETE("/:id", handler.Delete)
}
