package models

import (
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(event *Event) error {
	return r.db.Create(event).Error
}

type ClubRepository struct {
	db *gorm.DB
}

func NewClubRepository(db *gorm.DB) *ClubRepository {
	return &ClubRepository{db: db}
}

func (r *ClubRepository) AddMember(clubID, userID uint) error {
	var club Club
	if err := r.db.First(&club, clubID).Error; err != nil {
		return err
	}

	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	return r.db.Model(&club).Association("Members").Append(&user)
}

func (r *ClubRepository) RemoveMember(clubID, userID uint) error {
	var club Club
	if err := r.db.First(&club, clubID).Error; err != nil {
		return err
	}

	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}

	return r.db.Model(&club).Association("Members").Delete(&user)
}
