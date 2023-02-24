package repository

import (
	"database/sql"

	"log"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/coordinate"
	location "github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/location"
	customerror "github.com/kokiebisu/rental-storage/service-space/internal/error"
	"github.com/kokiebisu/rental-storage/service-space/internal/helper"
)

type SpaceRepository struct {
	db *sql.DB
}

func NewSpaceRepository(db *sql.DB) *SpaceRepository {
	return &SpaceRepository{
		db: db,
	}
}

func (r *SpaceRepository) Setup() *customerror.CustomError {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS spaces (
			id SERIAL NOT NULL PRIMARY KEY, 
			uid UUID NOT NULL,
			title VARCHAR(64) NOT NULL,
			lender_id UUID NOT NULL, 
			description TEXT NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP
	  	)
	`)
	if err != nil {
		log.Fatalf(err.Error())
		return customerror.ErrorHandler.CreateTableError("spaces", err)
		// ROLLBACK
	}
	_, err = r.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS images (
		  id SERIAL PRIMARY KEY,
		  space_id INTEGER REFERENCES spaces(id) ON DELETE CASCADE,
		  url TEXT NOT NULL
		)
	  `,
	)
	if err != nil {
		log.Fatalf(err.Error())
		// ROLLBACK
		return customerror.ErrorHandler.CreateTableError("images", err)
	}
	_, err = r.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS locations (
		  id SERIAL PRIMARY KEY,
		  space_id INTEGER UNIQUE REFERENCES spaces(id),
		  address VARCHAR(255),
		  city VARCHAR(100),
		  country VARCHAR(50),
		  country_code VARCHAR(3),
		  phone_number VARCHAR(20),
		  province VARCHAR(50),
		  province_code VARCHAR(4),
		  zip VARCHAR(20),
		  coordinate point
		)
	  `,
	)
	if err != nil {
		log.Fatalf(err.Error())
		// ROLLBACK
		return customerror.ErrorHandler.CreateTableError("locations", err)
	}
	return nil
}

func (r *SpaceRepository) Save(space space.Entity) (string, *customerror.CustomError) {
	var lastInsertedId int64
	row := r.db.QueryRow(
		`
          INSERT INTO spaces (
          uid, title, lender_id, description, created_at, updated_at
          ) VALUES ($1, $2, $3, $4, $5, $6)
		  RETURNING id
		`,
		space.UId,
		space.Title,
		space.LenderId,
		space.Description,
		space.CreatedAt,
		space.UpdatedAt,
	)
	err := row.Scan(&lastInsertedId)
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	for _, url := range space.ImageUrls {
		_, err := r.db.Exec(
			`INSERT INTO images (url, space_id) VALUES ($1, $2)`,
			url,
			lastInsertedId,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	_, err = r.db.Exec(
		`
          INSERT INTO locations (
		  space_id, address, city, country, country_code, phone_number, province, province_code, zip, coordinate
          ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, point($10, $11))
		  RETURNING id
		`,
		lastInsertedId,
		space.Location.Address,
		space.Location.City,
		space.Location.Country,
		space.Location.CountryCode,
		space.Location.Phone,
		space.Location.Province,
		space.Location.ProvinceCode,
		space.Location.Zip,
		space.Location.Coordinate.Latitude,
		space.Location.Coordinate.Longitude,
	)
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	return space.UId, nil
}

func (r *SpaceRepository) Delete(uid string) (string, *customerror.CustomError) {
	_, err := r.db.Exec(`DELETE FROM spaces WHERE uid = $1`)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteSpaceRowError("space", err)
	}
	return uid, nil
}

func (r *SpaceRepository) FindOneById(uid string) (space.Entity, *customerror.CustomError) {
	rows, err := r.db.Query(
		`
          SELECT spaces.*, images.url, locations.*, FROM spaces
          LEFT JOIN images ON spaces.id = images.space_id
		  LEFT JOIN locations ON spaces.id = locations.space_id
          WHERE spaces.uid = $1
        `,
		uid,
	)
	if err != nil {
		return space.Entity{}, customerror.ErrorHandler.FindSpacesRowError(err)
	}
	var id string
	var title string
	var lenderId string
	var description string
	var createdAt string
	var updatedAt string
	var imageUrls []string
	var address string
	var city string
	var country string
	var countryCode string
	var phoneNumber string
	var province string
	var provinceCode string
	var zip string
	var point helper.Point

	for rows.Next() {
		var uid string
		var imageUrl string
		err = rows.Scan(&id, &uid, &title, &lenderId, &description, &createdAt, &updatedAt, &imageUrl, &address, &city, &country, &countryCode, &phoneNumber, &province, &provinceCode, &zip, &point)
		if err != nil {
			return space.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		imageUrls = append(imageUrls, imageUrl)
	}
	l := space.Raw{
		UId:      uid,
		Title:    title,
		LenderId: lenderId,
		Location: location.Raw{
			Address:      address,
			City:         city,
			Country:      country,
			CountryCode:  countryCode,
			Phone:        phoneNumber,
			Province:     province,
			ProvinceCode: provinceCode,
			Zip:          zip,
			Coordinate: coordinate.Raw{
				Latitude:  point.X,
				Longitude: point.Y,
			},
		},
		Description: description,
		ImageUrls:   imageUrls,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}.ToEntity()
	return l, nil
}

func (r *SpaceRepository) FindManyByUserId(userId string) ([]space.Entity, *customerror.CustomError) {
	rows, err := r.db.Query(
		`
		SELECT spaces.*, images.url, locations.* FROM spaces 
		LEFT JOIN images ON spaces.id = spaces.space_id
		LEFT JOIN locations ON spaces.id = locations.space_id
		WHERE spaces.lender_id = $1
		`,
		userId,
	)
	if err != nil {
		return []space.Entity{}, customerror.ErrorHandler.FindSpacesRowError(err)
	}
	spacesMap := map[string]space.Entity{}
	for rows.Next() {
		var id string
		var uid string
		var title string
		var lenderId string
		var description string
		var createdAt string
		var updatedAt string
		var imageUrl string
		var address string
		var city string
		var country string
		var countryCode string
		var phoneNumber string
		var province string
		var provinceCode string
		var zip string
		var point helper.Point

		err := rows.Scan(&id, &uid, &title, &lenderId, &description, &createdAt, &updatedAt, &imageUrl, &address, &city, &country, &countryCode, &phoneNumber, &province, &provinceCode, &zip, &point)
		if err != nil {
			return []space.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		if entry, ok := spacesMap[uid]; ok {
			entry.ImageUrls = append(spacesMap[uid].ImageUrls, imageUrl)
		} else {
			l := space.Raw{
				UId:         uid,
				Title:       title,
				LenderId:    lenderId,
				Description: description,
				ImageUrls:   append([]string{}, imageUrl),
				Location: location.Raw{
					Address:      address,
					City:         city,
					Country:      country,
					CountryCode:  countryCode,
					Phone:        phoneNumber,
					Province:     province,
					ProvinceCode: provinceCode,
					Zip:          zip,
					Coordinate: coordinate.Raw{
						Latitude:  point.X,
						Longitude: point.Y,
					},
				},
				CreatedAt: createdAt,
				UpdatedAt: updatedAt,
			}.ToEntity()
			spacesMap[uid] = l
		}
	}
	spaces := []space.Entity{}
	for _, value := range spacesMap {
		spaces = append(spaces, value)
	}
	return spaces, nil
}
