package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(r *gin.Engine, db *gorm.DB) {
	user := r.Group("/users")
	UserRoutes(user, db)

	admin := r.Group("/admin")
	AdminRoutes(admin, db)
}
