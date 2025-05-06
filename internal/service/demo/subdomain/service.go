package subdomain

type Service struct {
	allEntities map[uint64]*Subdomain
}

func NewService() *Service {
	return &Service{
		allEntities: make(map[uint64]*Subdomain),
	}
}

func (s *Service) List() map[uint64]*Subdomain {
	return s.allEntities
}

func (s *Service) Get(entityID uint64) (*Subdomain, error) {
	return s.allEntities[entityID], nil
}

func (s *Service) New(title string) {
	s.allEntities[uint64(len(s.allEntities))] = &Subdomain{
		Title: title,
	}
}

func (s *Service) Delete(entityID uint64) bool {
	_, exists := s.allEntities[entityID]
	if exists {
		delete(s.allEntities, entityID)
	}
	return exists
}

func (s *Service) Edit(entityID uint64, title string) bool {
	entity, exists := s.allEntities[entityID]
	if exists {
		entity.Title = title
	}
	return exists
}
