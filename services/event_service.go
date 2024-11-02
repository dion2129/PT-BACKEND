// services/event_service.go
package services

import "api-test/models"

type EventService struct {
	repo *models.EventRepository
}

func NewEventService(repo *models.EventRepository) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) CreateEvent(event *models.Event) error {
	return s.repo.CreateEvent(event)
}
