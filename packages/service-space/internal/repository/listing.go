package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/amount"
	"github.com/kokiebisu/rental-storage/service-space/internal/core/domain/space/fee"
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
		CREATE TABLE IF NOT EXISTS space (
			id SERIAL NOT NULL PRIMARY KEY, 
			uid VARCHAR(64) NOT NULL,
			title VARCHAR(64),
			lender_id VARCHAR(64), 
			street_address VARCHAR(100) NOT NULL, 
			coordinate point
	  	)
	`)
	if err != nil {
		log.Fatalf(err.Error())
		return customerror.ErrorHandler.CreateTableError("space", err)
		// ROLLBACK
	}
	_, err = r.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS images_space (
		  url TEXT NOT NULL,
		  space_id INT NOT NULL,
		  PRIMARY KEY (url, space_id),
		  CONSTRAINT fk_space
			FOREIGN KEY(space_id) 
			  REFERENCES space(id)
			  ON DELETE CASCADE
		)
	  `,
	)
	if err != nil {
		log.Fatalf(err.Error())
		// ROLLBACK
		return customerror.ErrorHandler.CreateTableError("images_space", err)
	}
	_, err = r.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS fees_space (
		  amount INT NOT NULL,
		  currency VARCHAR(6) NOT NULL,
		  type VARCHAR(10) NOT NULL,
		  space_id INT NOT NULL,
		  PRIMARY KEY(space_id, type),
		  CONSTRAINT fk_space
			FOREIGN KEY(space_id) 
			  REFERENCES space(id)
			  ON DELETE CASCADE
		)
	  `,
	)
	if err != nil {
		log.Fatalf(err.Error())
		// ROLLBACK
		return customerror.ErrorHandler.CreateTableError("fees_space", err)
	}
	return nil
}

func (r *SpaceRepository) Save(space space.Entity) (string, *customerror.CustomError) {
	var lastInsertedId int64
	row := r.db.QueryRow(
		`
          INSERT INTO space (
          uid, title, lender_id, street_address, coordinate
          ) VALUES ($1, $2, $3, $4, point($5, $6))
		  RETURNING id
		`,
		space.UId,
		space.Title,
		space.LenderId,
		space.StreetAddress.Value,
		space.Latitude.Value,
		space.Longitude.Value,
	)
	err := row.Scan(&lastInsertedId)
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	for _, url := range space.ImageUrls {
		_, err := r.db.Exec(
			`INSERT INTO images_space (url, space_id) VALUES ($1, $2)`,
			url,
			lastInsertedId,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = r.db.Exec(
		`INSERT INTO fees_space (amount, currency, type, space_id) VALUES ($1, $2, $3, $4)`,
		space.Fee.Amount.Value,
		space.Fee.Amount.Currency,
		space.Fee.Type,
		lastInsertedId,
	)
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	return space.UId, nil
}

func (r *SpaceRepository) Delete(uid string) (string, *customerror.CustomError) {
	var removedSpaceId int32
	result := r.db.QueryRow(`DELETE FROM space WHERE uid = $1 RETURNING id`, uid)
	err := result.Scan(&removedSpaceId)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteSpaceRowError("space", err)
	}
	_, err = r.db.Exec(`DELETE FROM images_space WHERE space_id = $1`, removedSpaceId)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteSpaceRowError("images_space", err)
	}
	_, err = r.db.Exec(`DELETE FROM fees_space WHERE space_id = $1 RETURNING *`, removedSpaceId)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteSpaceRowError("fees_space", err)
	}
	return uid, nil
}

func (r *SpaceRepository) FindOneById(uid string) (space.Entity, *customerror.CustomError) {
	rows, err := r.db.Query(
		`
          SELECT space.*, images_space.url, fees_space.amount, fees_space.currency, fees_space.type FROM space 
          LEFT JOIN images_space ON space.id = images_space.space_id
          LEFT JOIN fees_space ON space.id = fees_space.space_id
          WHERE space.uid = $1
        `,
		uid,
	)
	if err != nil {
		return space.Entity{}, customerror.ErrorHandler.FindSpacesRowError(err)
	}
	var id string
	var title string
	var lenderId string
	var streetAddress string
	var coordinate helper.Point
	var feeAmount int64
	var feeCurrency string
	var feeType string
	var imageUrls []string

	for rows.Next() {
		var uid string
		var imageUrl string
		err = rows.Scan(&id, &uid, &title, &lenderId, &streetAddress, &coordinate, &imageUrl, &feeAmount, &feeCurrency, &feeType)
		if err != nil {
			return space.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		imageUrls = append(imageUrls, imageUrl)
	}
	l := space.Raw{
		UId:           uid,
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: streetAddress,
		Latitude:      coordinate.X,
		Longitude:     coordinate.Y,
		ImageUrls:     imageUrls,
		Fee: fee.Raw{
			Amount: amount.Raw{
				Value:    feeAmount,
				Currency: feeCurrency,
			},
			Type: feeType,
		},
	}.ToEntity()
	return l, nil
}

// deprecating this for now
// not sure if this is valid
func (r *SpaceRepository) FindManyByLatLng(latitude float64, longitude float64, distance int32) ([]space.Entity, *customerror.CustomError) {
	rows, err := r.db.Query(
		`
			SELECT * FROM (
				SELECT space.*, 
						( 3959 * acos( cos( radians($1) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians($2) ) + sin( radians(1) ) * sin( radians( latitude ) ) ) ) 
						AS distance, images_space.url, fees_space.amount, fees_space.currency, fees_space.type FROM space 
						LEFT JOIN images_space ON space.id = images_space.space_id
						LEFT JOIN fees_space ON space.id = fees_space.space_id
						GROUP BY id, uid, images_space.url, fees_space.amount, fees_space.currency, fees_space.type
			) 
			x GROUP BY x.id, x.uid, x.title, x.lender_id, x.street_address, x.coordinate, x.url, x.amount, x.currency, x.type, x.distance 
			HAVING x.distance < $3 ORDER BY x.distance LIMIT 10
		`,
		latitude, longitude, distance,
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
		var streetAddress string
		var coordinate helper.Point
		var distance float32
		var imageUrl string
		var feeAmount int64
		var feeCurrency string
		var feeType string
		err := rows.Scan(&id, &uid, &title, &lenderId, &streetAddress, &coordinate, &distance, &imageUrl, &feeAmount, &feeCurrency, &feeType)
		if err != nil {
			return []space.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		if entry, ok := spacesMap[uid]; ok {
			entry.ImageUrls = append(spacesMap[uid].ImageUrls, imageUrl)
		} else {
			l := space.Raw{
				UId:           uid,
				Title:         title,
				LenderId:      lenderId,
				StreetAddress: streetAddress,
				Latitude:      latitude,
				Longitude:     longitude,
				ImageUrls:     append([]string{}, imageUrl),
				Fee: fee.Raw{
					Amount: amount.Raw{
						Value:    feeAmount,
						Currency: feeCurrency,
					},
					Type: feeType,
				},
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

func (r *SpaceRepository) FindManyByUserId(userId string) ([]space.Entity, *customerror.CustomError) {
	fmt.Println("ENTERED1")
	rows, err := r.db.Query(
		`
		SELECT space.*, images_space.url, fees_space.amount, fees_space.currency, fees_space.type FROM space 
		LEFT JOIN images_space ON space.id = images_space.space_id
		LEFT JOIN fees_space ON space.id = fees_space.space_id
		WHERE space.lender_id = $1
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
		var streetAddress string
		var coordinate helper.Point
		var imageUrl string
		var feeAmount int64
		var feeCurrency string
		var feeType string
		err := rows.Scan(&id, &uid, &title, &lenderId, &streetAddress, &coordinate, &imageUrl, &feeAmount, &feeCurrency, &feeType)
		if err != nil {
			return []space.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		if entry, ok := spacesMap[uid]; ok {
			entry.ImageUrls = append(spacesMap[uid].ImageUrls, imageUrl)
		} else {
			l := space.Raw{
				UId:           uid,
				Title:         title,
				LenderId:      lenderId,
				StreetAddress: streetAddress,
				Latitude:      coordinate.X,
				Longitude:     coordinate.Y,
				ImageUrls:     append([]string{}, imageUrl),
				Fee: fee.Raw{
					Amount: amount.Raw{
						Value:    feeAmount,
						Currency: feeCurrency,
					},
					Type: feeType,
				},
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
