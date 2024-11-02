// handler/club_handler.go
package handler

import (
	"api-test/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClubHandler struct {
	clubService *services.ClubService
}

func NewClubHandler(clubService *services.ClubService) *ClubHandler {
	return &ClubHandler{clubService: clubService}
}

func (h *ClubHandler) AddMember(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id" binding:"required"`
		ClubID uint `json:"club_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.clubService.AddMember(input.ClubID, input.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member added successfully"})
}

func (h *ClubHandler) RemoveMember(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id" binding:"required"`
		ClubID uint `json:"club_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.clubService.RemoveMember(input.ClubID, input.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "member removed successfully"})
}
