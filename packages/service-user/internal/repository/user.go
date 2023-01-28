package repository

import (
	"database/sql"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	errors "github.com/kokiebisu/rental-storage/service-user/internal/error"
	_ "github.com/lib/pq"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Setup() *errors.CustomError {
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
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func (r *UserRepository) Save(user domain.User) (string, *errors.CustomError) {
	_, err := r.db.Exec(`
		INSERT INTO user_account (uid, email_address, password, first_name, last_name, created_at, updated_at) 
        VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, user.Uid, user.EmailAddress.Value, user.Password, user.FirstName.Value, user.LastName.Value, user.CreatedAt.Format("2006-01-02"), user.UpdatedAt.Format("2006-01-02"))
	if err != nil {
		return "", errors.ErrorHandler.InternalServerError()
	}
	return user.Uid, nil
}

func (r *UserRepository) Delete(uid string) *errors.CustomError {
	_, err := r.db.Exec(`
		DELETE FROM user_account WHERE uid = $1
	`, uid)
	if err != nil {
		return errors.ErrorHandler.InternalServerError()
	}
	return nil
}

func (r *UserRepository) FindOneById(uid string) (domain.User, *errors.CustomError) {

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
		return domain.User{}, errors.ErrorHandler.InternalServerError()
	}
	user := &domain.UserRaw{
		Uid:          uid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Password:     password,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
	return user.ToEntity(), nil
}

func (r *UserRepository) FindOneByEmail(emailAddress string) (domain.User, *errors.CustomError) {
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

	if err != nil {
		return domain.User{}, errors.ErrorHandler.InternalServerError()
	}
	user := &domain.UserRaw{
		Uid:          uid,
		FirstName:    firstName,
		LastName:     lastName,
		EmailAddress: emailAddress,
		Password:     password,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
	return user.ToEntity(), nil
}
