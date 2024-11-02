package routes

import (
	"api-test/handler"
	"api-test/middleware"
	"api-test/models"
	"api-test/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AdminRoutes(group *gin.RouterGroup, db *gorm.DB) {
	eventRepo := models.NewEventRepository(db)
	eventService := services.NewEventService(eventRepo)
	eventHandler := handler.NewEventHandler(eventService)

	clubRepo := models.NewClubRepository(db)
	clubService := services.NewClubService(clubRepo)
	clubHandler := handler.NewClubHandler(clubService)

	adminGroup := group.Use(middleware.AdminMiddleware())
	adminGroup.POST("/events", eventHandler.CreateEvent)
	adminGroup.POST("/clubs/members", clubHandler.AddMember)
	adminGroup.DELETE("/clubs/members", clubHandler.RemoveMember)
}