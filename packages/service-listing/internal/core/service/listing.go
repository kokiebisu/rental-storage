package service

import (
	"log"

	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

type ListingService struct {
	ListingRepository port.ListingRepository
	ListingFactory    port.ListingFactory
}

func NewListingService(listingRepository port.ListingRepository, listingFactory port.ListingFactory) *ListingService {
	return &ListingService{
		ListingRepository: listingRepository,
		ListingFactory:    listingFactory,
	}
}

func (s *ListingService) FindListingsWithinLatLng(latitude float64, longitude float64, distance int32) ([]listing.DTO, *errors.CustomError) {
	listings, err := s.ListingRepository.FindManyByLatLng(latitude, longitude, distance)
	if err != nil {
		return []listing.DTO{}, err
	}
	listingDTOs := []listing.DTO{}
	for _, l := range listings {
		listingDTO, err := l.ToDTO()
		if err != nil {
			log.Fatalf(err.Error())
			return []listing.DTO{}, err
		}
		listingDTOs = append(listingDTOs, listingDTO)
	}
	return listingDTOs, nil
}

func (s *ListingService) FindListingById(uid string) (listing.DTO, *errors.CustomError) {
	l, err := s.ListingRepository.FindOneById(uid)
	if err != nil {
		return listing.DTO{}, err
	}
	listingDTO, err := l.ToDTO()
	if err != nil {
		return listing.DTO{}, err
	}
	return listingDTO, nil
}

func (s *ListingService) CreateListing(lenderId string, streetAddress string, latitude float64, longitude float64, imageUrls []string, title string, feeAmount int32, feeCurrency amount.CurrencyType, feeType fee.RentalFeeType) (string, *errors.CustomError) {
	entity, err := s.ListingFactory.New(title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, int64(feeAmount), string(feeType))
	if err != nil {
		return "", err
	}
	result, err := s.ListingRepository.Save(entity)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (s *ListingService) RemoveListingById(uid string) *errors.CustomError {
	err := s.ListingRepository.Delete(uid)
	return err
}
