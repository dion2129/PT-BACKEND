package handler

import (
	"api-test/helpers"
	"api-test/models"
	"api-test/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

// NewUserHandler adalah konstruktor untuk UserHandler
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Fungsi Register untuk pendaftaran user
func (h *UserHandler) Register(c *gin.Context) {
	var user models.User
	fmt.Println("baris 1")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := helpers.HashPass(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		fmt.Println("baris 2")

		return
	}
	user.Password = hashedPassword

	newUser, err := h.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		fmt.Println("baris 3")

		return
	}

	c.JSON(http.StatusOK, gin.H{"user": newUser})
	fmt.Println("baris 4")

}

// Fungsi Login untuk autentikasi user
func (h *UserHandler) Login(c *gin.Context) {
	var loginDTO models.LoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUserByEmail(loginDTO.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	match, err := helpers.ComparePass([]byte(user.Password), []byte(loginDTO.Password))
	if err != nil || !match {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := helpers.CreateTokenJWT(int(user.ID), "user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Fungsi GetProfile untuk mengambil profil user berdasarkan ID
func (h *UserHandler) GetProfile(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Fungsi GetAllUser untuk mengambil semua user
func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Fungsi Delete untuk menghapus user berdasarkan ID
func (h *UserHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
