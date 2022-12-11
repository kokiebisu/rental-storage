package repository

import (
	"github.com/kokiebisu/rental-storage/service-booking/internal/adapter/db"
	"github.com/kokiebisu/rental-storage/service-booking/internal/core/domain/booking"
)

type BookingRepository struct {
	db *db.NoSQLClient
}

func NewBookingRepository(db *db.NoSQLClient) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

func (r *BookingRepository) Save(booking booking.Entity) error {
	err := r.db.Save(booking)
	return err
}

func (r *BookingRepository) Delete(id string) error {
	err := r.db.Delete(id)
	return err
}

func (r *BookingRepository) FindOneById(id string) (booking.Entity, error) {
	entity, err := r.db.FindById(id)
	if err != nil {
		return booking.Entity{}, err
	}
	return entity, nil
}

func (r *BookingRepository) FindManyByUserId(userId string) ([]booking.Entity, error) {
	entities, err := r.db.FindManyByUserId(userId)
	if err != nil {
		return []booking.Entity{}, err
	}
	return entities, nil
}
