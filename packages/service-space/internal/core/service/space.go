package service

import (
	"log"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/port"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
)

type SpaceService struct {
	SpaceRepository port.SpaceRepository
	SpaceFactory    port.SpaceFactory
}

func NewSpaceService(spaceRepository port.SpaceRepository, spaceFactory port.SpaceFactory) *SpaceService {
	return &SpaceService{
		SpaceRepository: spaceRepository,
		SpaceFactory:    spaceFactory,
	}
}

func (s *SpaceService) FindSpacesByUserId(userId string) ([]space.DTO, *customerror.CustomError) {
	ls, err := s.SpaceRepository.FindManyByUserId(userId)
	if err != nil {
		return []space.DTO{}, err
	}
	spaceDTOs := []space.DTO{}
	for _, l := range ls {
		spaceDTO, err := l.ToDTO()
		if err != nil {
			log.Fatalf(err.Error())
			return []space.DTO{}, err
		}
		spaceDTOs = append(spaceDTOs, spaceDTO)
	}
	return spaceDTOs, nil
}

func (s *SpaceService) FindSpacesWithinLatLng(latitude float64, longitude float64, distance int32) ([]space.DTO, *customerror.CustomError) {
	spaces, err := s.SpaceRepository.FindManyByLatLng(latitude, longitude, distance)
	if err != nil {
		return []space.DTO{}, err
	}
	spaceDTOs := []space.DTO{}
	for _, l := range spaces {
		spaceDTO, err := l.ToDTO()
		if err != nil {
			log.Fatalf(err.Error())
			return []space.DTO{}, err
		}
		spaceDTOs = append(spaceDTOs, spaceDTO)
	}
	return spaceDTOs, nil
}

func (s *SpaceService) FindSpaceById(uid string) (space.DTO, *customerror.CustomError) {
	l, err := s.SpaceRepository.FindOneById(uid)
	if err != nil {
		return space.DTO{}, err
	}
	spaceDTO, err := l.ToDTO()
	if err != nil {
		return space.DTO{}, err
	}
	return spaceDTO, nil
}

func (s *SpaceService) CreateSpace(lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, title string) (string, *customerror.CustomError) {
	entity, err := s.SpaceFactory.New(title, lenderId, streetAddress, latitude, longitude, imageUrls)
	if err != nil {
		return "", err
	}
	result, err := s.SpaceRepository.Save(entity)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *SpaceService) RemoveSpaceById(uid string) (string, *customerror.CustomError) {
	return s.SpaceRepository.Delete(uid)
}
