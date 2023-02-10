package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/amount"
	"github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing/fee"
	customerror "github.com/kokiebisu/rental-storage/service-listing/internal/error"
	"github.com/kokiebisu/rental-storage/service-listing/internal/helper"
)

type ListingRepository struct {
	db *sql.DB
}

func NewListingRepository(db *sql.DB) *ListingRepository {
	return &ListingRepository{
		db: db,
	}
}

func (r *ListingRepository) Setup() *customerror.CustomError {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS listing (
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
		return customerror.ErrorHandler.CreateTableError("listing", err)
		// ROLLBACK
	}
	_, err = r.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS images_listing (
		  url TEXT NOT NULL,
		  listing_id INT NOT NULL,
		  PRIMARY KEY (url, listing_id),
		  CONSTRAINT fk_listing
			FOREIGN KEY(listing_id) 
			  REFERENCES listing(id)
			  ON DELETE CASCADE
		)
	  `,
	)
	if err != nil {
		log.Fatalf(err.Error())
		// ROLLBACK
		return customerror.ErrorHandler.CreateTableError("images_listing", err)
	}
	_, err = r.db.Exec(
		`
		CREATE TABLE IF NOT EXISTS fees_listing (
		  amount INT NOT NULL,
		  currency VARCHAR(6) NOT NULL,
		  type VARCHAR(10) NOT NULL,
		  listing_id INT NOT NULL,
		  PRIMARY KEY(listing_id, type),
		  CONSTRAINT fk_listing
			FOREIGN KEY(listing_id) 
			  REFERENCES listing(id)
			  ON DELETE CASCADE
		)
	  `,
	)
	if err != nil {
		log.Fatalf(err.Error())
		// ROLLBACK
		return customerror.ErrorHandler.CreateTableError("fees_listing", err)
	}
	return nil
}

func (r *ListingRepository) Save(listing listing.Entity) (string, *customerror.CustomError) {
	var lastInsertedId int64
	row := r.db.QueryRow(
		`
          INSERT INTO listing (
          uid, title, lender_id, street_address, coordinate
          ) VALUES ($1, $2, $3, $4, point($5, $6))
		  RETURNING id
		`,
		listing.UId,
		listing.Title,
		listing.LenderId,
		listing.StreetAddress.Value,
		listing.Latitude.Value,
		listing.Longitude.Value,
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
	for _, url := range listing.ImageUrls {
		_, err := r.db.Exec(
			`INSERT INTO images_listing (url, listing_id) VALUES ($1, $2)`,
			url,
			lastInsertedId,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
	_, err = r.db.Exec(
		`INSERT INTO fees_listing (amount, currency, type, listing_id) VALUES ($1, $2, $3, $4)`,
		listing.Fee.Amount.Value,
		listing.Fee.Amount.Currency,
		listing.Fee.Type,
		lastInsertedId,
	)
	if err != nil {
		log.Fatal(err.Error())
		// ROLLBACK
	}
	return listing.UId, nil
}

func (r *ListingRepository) Delete(uid string) (string, *customerror.CustomError) {
	var removedListingId int32
	result := r.db.QueryRow(`DELETE FROM listing WHERE uid = $1 RETURNING id`, uid)
	err := result.Scan(&removedListingId)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteListingRowError("listing", err)
	}
	_, err = r.db.Exec(`DELETE FROM images_listing WHERE listing_id = $1`, removedListingId)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteListingRowError("images_listing", err)
	}
	_, err = r.db.Exec(`DELETE FROM fees_listing WHERE listing_id = $1 RETURNING *`, removedListingId)
	if err != nil {
		return "", customerror.ErrorHandler.DeleteListingRowError("fees_listing", err)
	}
	return uid, nil
}

func (r *ListingRepository) FindOneById(uid string) (listing.Entity, *customerror.CustomError) {
	rows, err := r.db.Query(
		`
          SELECT listing.*, images_listing.url, fees_listing.amount, fees_listing.currency, fees_listing.type FROM listing 
          LEFT JOIN images_listing ON listing.id = images_listing.listing_id
          LEFT JOIN fees_listing ON listing.id = fees_listing.listing_id
          WHERE listing.uid = $1
        `,
		uid,
	)
	if err != nil {
		return listing.Entity{}, customerror.ErrorHandler.FindListingsRowError(err)
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
			return listing.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		imageUrls = append(imageUrls, imageUrl)
	}
	l := listing.Raw{
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
func (r *ListingRepository) FindManyByLatLng(latitude float64, longitude float64, distance int32) ([]listing.Entity, *customerror.CustomError) {
	rows, err := r.db.Query(
		`
			SELECT * FROM (
				SELECT listing.*, 
						( 3959 * acos( cos( radians($1) ) * cos( radians( latitude ) ) * cos( radians( longitude ) - radians($2) ) + sin( radians(1) ) * sin( radians( latitude ) ) ) ) 
						AS distance, images_listing.url, fees_listing.amount, fees_listing.currency, fees_listing.type FROM listing 
						LEFT JOIN images_listing ON listing.id = images_listing.listing_id
						LEFT JOIN fees_listing ON listing.id = fees_listing.listing_id
						GROUP BY id, uid, images_listing.url, fees_listing.amount, fees_listing.currency, fees_listing.type
			) 
			x GROUP BY x.id, x.uid, x.title, x.lender_id, x.street_address, x.coordinate, x.url, x.amount, x.currency, x.type, x.distance 
			HAVING x.distance < $3 ORDER BY x.distance LIMIT 10
		`,
		latitude, longitude, distance,
	)
	if err != nil {
		return []listing.Entity{}, customerror.ErrorHandler.FindListingsRowError(err)
	}
	listingsMap := map[string]listing.Entity{}
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
			return []listing.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		if entry, ok := listingsMap[uid]; ok {
			entry.ImageUrls = append(listingsMap[uid].ImageUrls, imageUrl)
		} else {
			l := listing.Raw{
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
			listingsMap[uid] = l
		}
	}
	listings := []listing.Entity{}
	for _, value := range listingsMap {
		listings = append(listings, value)
	}
	return listings, nil
}

func (r *ListingRepository) FindManyByUserId(userId string) ([]listing.Entity, *customerror.CustomError) {
	fmt.Println("ENTERED1")
	rows, err := r.db.Query(
		`
		SELECT listing.*, images_listing.url, fees_listing.amount, fees_listing.currency, fees_listing.type FROM listing 
		LEFT JOIN images_listing ON listing.id = images_listing.listing_id
		LEFT JOIN fees_listing ON listing.id = fees_listing.listing_id
		WHERE listing.lender_id = $1
		`,
		userId,
	)
	if err != nil {
		return []listing.Entity{}, customerror.ErrorHandler.FindListingsRowError(err)
	}
	listingsMap := map[string]listing.Entity{}
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
			return []listing.Entity{}, customerror.ErrorHandler.ScanRowError(err)
		}
		if entry, ok := listingsMap[uid]; ok {
			entry.ImageUrls = append(listingsMap[uid].ImageUrls, imageUrl)
		} else {
			l := listing.Raw{
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
			listingsMap[uid] = l
		}
	}
	listings := []listing.Entity{}
	for _, value := range listingsMap {
		listings = append(listings, value)
	}
	return listings, nil
}
