package repository

import (
	"database/sql"

	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/item"
	"github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	customerror "github.com/kokiebisu/rental-storage/service-user/internal/error"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) Setup() *customerror.CustomError {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS user_account (
			id SERIAL NOT NULL PRIMARY KEY, 
			uid VARCHAR(64) NOT NULL, 
			first_name VARCHAR(32) NOT NULL, 
			last_name VARCHAR(32) NOT NULL, 
			email_address VARCHAR(64) NOT NULL UNIQUE, 
			password VARCHAR(64) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP
		)
	`)
	if err != nil {
		return customerror.ErrorHandler.CreateTableError("user_account", err)
	}
	return nil
}

func (r *UserRepository) Save(u user.Entity) (string, *customerror.CustomError) {
	_, err := r.db.Exec(`
		INSERT INTO user_account (uid, email_address, password, first_name, last_name, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, u.UId, u.EmailAddress, u.Password, u.FirstName, u.LastName, u.CreatedAt.Format("2006-01-02"), u.UpdatedAt.Format("2006-01-02"))
	if err != nil {
		return u.UId, customerror.ErrorHandler.InsertRowError("user_account", err)
	}
	return u.UId, nil
}

func (r *UserRepository) Delete(uid string) (string, *customerror.CustomError) {
	_, err := r.db.Exec(`
		DELETE FROM user_account WHERE uid = $1
	`, uid)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteRowError("uid", err)
	}
	return uid, nil
}

func (r *UserRepository) FindById(uid string) (user.Entity, *customerror.CustomError) {

	var id int64
	var emailAddress string
	var firstName string
	var lastName string
	var password string
	var createdAt string
	var updatedAt string

	row := r.db.QueryRow(`
        SELECT * FROM user_account 
        WHERE uid = $1
    `, uid)

	err := row.Scan(&id, &uid, &firstName, &lastName, &emailAddress, &password, &createdAt, &updatedAt)
	if err != nil {
		return user.Entity{}, customerror.ErrorHandler.ScanRowError(err)
	}
	u := &user.Raw{
		UId:          uid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Password:     password,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
	return u.ToEntity(), nil
}

func (r *UserRepository) FindByEmail(emailAddress string) (user.Entity, *customerror.CustomError) {
	var id int64
	var uid string
	var firstName string
	var lastName string
	var password string
	var createdAt string
	var updatedAt string
	row := r.db.QueryRow(`
		SELECT * FROM user_account
		WHERE email_address = $1
  	`, emailAddress)
	err := row.Scan(&id, &uid, &firstName, &lastName, &emailAddress, &password, &createdAt, &updatedAt)
	u := &user.Raw{
		UId:          uid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Password:     password,
		Items:        []item.Raw{},
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
	if err != nil {
		return user.Entity{}, customerror.ErrorHandler.ScanRowError(err)
	}
	return u.ToEntity(), nil
}
