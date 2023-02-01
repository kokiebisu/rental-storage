package service

import (
	"log"

	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/port"
	errors "github.com/kokiebisu/rental-storage/service-listing/internal/error"
)

var (
	listingFactory = &listing.Factory{}
)

type ListingService struct {
	listingRepository port.ListingRepository
}

func NewListingService(listingRepository port.ListingRepository) *ListingService {
	return &ListingService{
		listingRepository: listingRepository,
	}
}

func (s *ListingService) FindListingsWithinLatLng(latitude float32, longitude float32, distance int32) ([]listing.DTO, *errors.CustomError) {
	listings, err := s.listingRepository.FindManyByLatLng(latitude, longitude, distance)
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
	l, err := s.listingRepository.FindOneById(uid)
	if err != nil {
		return listing.DTO{}, err
	}
	listingDTO, err := l.ToDTO()
	if err != nil {
		return listing.DTO{}, err
	}
	return listingDTO, nil
}

func (s *ListingService) CreateListing(lenderId string, streetAddress string, latitude float32, longitude float32, imageUrls []string, title string, feeAmount int32, feeCurrency amount.CurrencyType, feeType fee.RentalFeeType) (string, *errors.CustomError) {
	entity, err := listingFactory.New(title, lenderId, streetAddress, latitude, longitude, imageUrls, feeCurrency, int64(feeAmount), string(feeType))
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
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}
