package repository

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
	errors "github.com/kokiebisu/rental-storage/service-booking/internal/error"
)

type BookingRepository struct {
	db *db.NoSQLClient
}

func NewBookingRepository(db *db.NoSQLClient) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) Save(booking booking.Entity) *errors.CustomError {
	err := r.db.Save(booking)
	return err
}

func (r *BookingRepository) Delete(id string) *errors.CustomError {
	err := r.db.Delete(id)
	return err
}

func (r *BookingRepository) FindOneById(id string) (booking.Entity, *errors.CustomError) {
	entity, err := r.db.FindById(id)
	if err != nil {
		return booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	return entity, nil
}

func (r *BookingRepository) FindManyByUserId(userId string) ([]booking.Entity, *errors.CustomError) {
	entities, err := r.db.FindManyByUserId(userId)
	if err != nil {
		return []booking.Entity{}, errors.ErrorHandler.InternalServerError()
	}
	return entities, nil
}
