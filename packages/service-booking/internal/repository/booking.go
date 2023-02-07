package repository

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	customerror "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingRepository struct {
	db *db.NoSQLClient
}

func NewBookingRepository(db *db.NoSQLClient) *BookingRepository {
	return &BookingRepository{
		db,
	}
}

func (r *BookingRepository) Save(booking booking.Entity) *customerror.CustomError {
	err := r.db.Save(booking)
	return err
}

func (r *BookingRepository) Delete(id string) *customerror.CustomError {
	err := r.db.Delete(id)
	return err
}

func (r *BookingRepository) FindOneById(id string) (booking.Entity, *customerror.CustomError) {
	entity, err := r.db.FindById(id)
	if err != nil {
		return booking.Entity{}, err
	}
	return entity, nil
}

func (r *BookingRepository) FindManyByUserId(userId string) ([]booking.Entity, *customerror.CustomError) {
	entities, err := r.db.FindManyByUserId(userId)
	if err != nil {
		return []booking.Entity{}, err
	}
	return entities, nil
}
