package service

import (
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceService struct {
	SpaceRepository port.SpaceRepository
}

func NewSpaceService(spaceRepository port.SpaceRepository) *SpaceService {
	return &SpaceService{
		SpaceRepository: spaceRepository,
	}
}

func (s *SpaceService) FindSpacesByUserId(userId string) ([]space.DTO, *customerror.CustomError) {
	ls, err := s.SpaceRepository.FindManyByUserId(userId)
	if err != nil {
		return []space.DTO{}, err
	}
	spaceDTOs := []space.DTO{}
	for _, l := range ls {
		spaceDTOs = append(spaceDTOs, l.ToDTO())
	}
	return spaceDTOs, nil
}

func (s *SpaceService) FindSpaceById(uid string) (space.DTO, *customerror.CustomError) {
	l, err := s.SpaceRepository.FindById(uid)
	if err != nil {
		return space.DTO{}, err
	}
	return l.ToDTO(), nil
}

func (s *SpaceService) CreateSpace(uid string, lenderId string, location location.DTO, imageUrls []string, title string, description string, createdAt string, updatedAt string) (string, *customerror.CustomError) {
	entity, err := space.New(uid, title, lenderId, location, imageUrls, description, createdAt, updatedAt)
	if err != nil {
		return "", err
	}
	result, err := s.SpaceRepository.Save(entity)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *SpaceService) DeleteSpaceById(uid string) (string, *customerror.CustomError) {
	return s.SpaceRepository.Delete(uid)
}
