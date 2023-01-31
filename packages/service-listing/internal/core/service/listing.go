package service

import (
	"log"

	domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingService struct {
	listingRepository port.ListingRepository
}

func NewListingService(listingRepository port.ListingRepository) *ListingService {
	return &ListingService{
		listingRepository: listingRepository,
	}
}

func (s *ListingService) FindListingsWithinLatLng(latitude float32, longitude float32, distance int32) ([]domain.ListingDTO, *errors.CustomError) {
	listings, err := s.listingRepository.FindManyByLatLng(latitude, longitude, distance)
	if err != nil {
		return []domain.ListingDTO{}, err
	}
	listingDTOs := []domain.ListingDTO{}
	for _, listing := range listings {
		listingDTO, err := listing.ToDTO()
		if err != nil {
			log.Fatalf(err.Error())
			return []domain.ListingDTO{}, err
		}
		listingDTOs = append(listingDTOs, listingDTO)
	}
	return listingDTOs, nil
}

func (s *ListingService) FindListingById(uid string) (domain.ListingDTO, *errors.CustomError) {
	listing, err := s.listingRepository.FindOneById(uid)
	if err != nil {
		return domain.ListingDTO{}, err
	}
	listingDTO, err := listing.ToDTO()
	if err != nil {
		return domain.ListingDTO{}, err
	}
	return listingDTO, nil
}

func (s *ListingService) CreateListing(lenderId string, streetAddress string, latitude float32, longitude float32, imageUrls []string, title string, feeAmount int32, feeCurrency domain.CurrencyType, feeType domain.RentalFeeType) (string, *errors.CustomError) {
	entity, err := domain.NewListing(title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, int64(feeAmount), string(feeType))
	if err != nil {
		return "", err
	}
	result, err := s.listingRepository.Save(entity)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *ListingService) RemoveListingById(uid string) *errors.CustomError {
	err := s.listingRepository.Delete(uid)
	if err != nil {
		return err
	}
	return nil
}
