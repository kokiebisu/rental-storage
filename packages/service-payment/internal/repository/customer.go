package repository

import (
	"database/sql"
	"service-payment/internal/core/domain"

	_ "github.com/lib/pq"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{
		db,
	}
}

func (r *CustomerRepository) SetupTables() error {
	_, err := r.db.Query(`
		CREATE TABLE IF NOT EXISTS customer (
			user_id VARCHAR(64) NOT NULL,
			customer_id VARCHAR(64) NOT NULL,
			provider_type VARCHAR(32) NOT NULL,
			PRIMARY KEY(user_id, provider_type)
		)
	`)
	if err != nil {
		return err
	}
	return nil
}

func (r *CustomerRepository) Save(pc domain.PaymentCustomer) (domain.PaymentCustomer, error) {
	_, err := r.db.Exec(`
		INSERT INTO customer (
			user_id,
			customer_id, 
			provider_type
		) VALUES ($1, $2, $3)
	`, pc.UserId, pc.CustomerId, pc.ProviderType)
	if err != nil {
		return domain.PaymentCustomer{}, err
	}
	return pc, nil
}
