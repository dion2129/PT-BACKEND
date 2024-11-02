package dto

import "api-test/models"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDetail struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserDetail(user models.User) UserDetail {
	return UserDetail{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func NewListUsers(users []models.User) []UserDetail {
	var listUser []UserDetail
	for _, user := range users {
		userDetail := NewUserDetail(user)
		listUser = append(listUser, userDetail)
	}
	return listUser
}
