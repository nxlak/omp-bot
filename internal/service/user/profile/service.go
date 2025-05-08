package profile

import (
	"errors"

	"github.com/ozonmp/omp-bot/internal/model/user"
)

type ProfileService interface {
	Describe(profileID uint64) (*user.Profile, error)
	List(cursor uint64, limit uint64) ([]user.Profile, error)
	Create(title string) (uint64, error)
	Update(profileID uint64, profile user.Profile) error
	Remove(profileID uint64) (bool, error)
}

type DummyProfileService struct{}

func NewDummyProfileService() *DummyProfileService {
	return &DummyProfileService{}
}

func (s *DummyProfileService) Describe(profileID uint64) (*user.Profile, error) {
	value, exists := data[profileID]
	if !exists {
		return nil, errors.New("user doesn't exist")
	}
	return value, nil
}

func (s *DummyProfileService) List(cursor uint64, limit uint64) ([]user.Profile, error) {
	resultItems := make([]user.Profile, 0)

	for idx, value := range data {
		if idx >= cursor && idx <= cursor+limit {
			resultItems = append(resultItems, *value)
		}
	}

	return resultItems, nil
}

func (s *DummyProfileService) Create(title string) (uint64, error) {
	profileId := uint64(len(data))
	data[profileId] = &user.Profile{ID: profileId, Title: title}
	return profileId, nil
}

func (s *DummyProfileService) Update(profileID uint64, title string) error {
	entity, exists := data[profileID]
	if !exists {
		return errors.New("Entity doesn't exist")
	}
	entity.Title = title
	return nil
}

func (s *DummyProfileService) Remove(profileID uint64) (bool, error) {
	_, exists := data[profileID]
	if exists {
		delete(data, profileID)
		return exists, nil
	}
	return exists, errors.New("Entity doesn't exist")
}
