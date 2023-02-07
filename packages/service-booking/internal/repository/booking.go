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
	return r.db.Save(booking)
}

func (r *BookingRepository) Delete(id string) *customerror.CustomError {
	return r.db.Delete(id)
}

func (r *BookingRepository) FindOneById(id string) (booking.Entity, *customerror.CustomError) {
	return r.db.FindById(id)
}

func (r *BookingRepository) FindManyByUserId(userId string) ([]booking.Entity, *customerror.CustomError) {
	return r.db.FindManyByUserId(userId)
}
