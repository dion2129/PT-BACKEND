// services/club_service.go
package services

import "api-test/models"

type ClubService struct {
	repo *models.ClubRepository
}

func NewClubService(repo *models.ClubRepository) *ClubService {
	return &ClubService{repo: repo}
}

func (s *ClubService) AddMember(clubID, userID uint) error {
	return s.repo.AddMember(clubID, userID)
}

func (s *ClubService) RemoveMember(clubID, userID uint) error {
	return s.repo.RemoveMember(clubID, userID)
}
