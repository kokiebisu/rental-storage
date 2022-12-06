package repository

import (
	"database/sql"
	"log"

	domain "github.com/kokiebisu/rental-storage/service-listing/internal/core/domain/listing"
)

type ListingRepository struct {
	db *sql.DB
}

func NewListingRepository(db *sql.DB) *ListingRepository {
	return &ListingRepository{
		db: db,
	}
}

func (r *ListingRepository) Setup() error {
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS listing (
			id SERIAL NOT NULL PRIMARY KEY, 
			uid VARCHAR(64) NOT NULL,
			title VARCHAR(64),
			lender_id VARCHAR(64), 
			street_address VARCHAR(100) NOT NULL, 
			latitude DECIMAL(16,14) NOT NULL, 
			longitude DECIMAL(16,14) NOT NULL
	  	)
	`)
	if err != nil {
		log.Fatalf(err.Error())
		return err
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
		return err
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
		return err
	}
	return nil
}

func (r *ListingRepository) Save(listing domain.Listing) (string, error) {
	var lastInsertedId int64
	row := r.db.QueryRow(
		`
          INSERT INTO listing (
          uid, lender_id, street_address, latitude, longitude
          ) VALUES ($1, $2, $3, $4, $5)
		  RETURNING id
		`,
		listing.Uid,
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
		r.db.Exec(
			`INSERT INTO images_listing (url, listing_id) VALUES ($1, $2)`,
			url,
			lastInsertedId,
		)
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
	return listing.Uid, nil
}

func (r *ListingRepository) Delete(uid string) error {
	var removedListingId int32
	result := r.db.QueryRow(`DELETE FROM listing WHERE uid = $1 RETURNING id`, uid)
	err := result.Scan(&removedListingId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`DELETE FROM images_listing WHERE listing_id = $1`, removedListingId)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(`DELETE FROM fees_listing WHERE listing_id = $1 RETURNING *`, removedListingId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ListingRepository) FindOneById(uid string) (domain.Listing, error) {
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
		log.Fatal(err.Error())
		return domain.Listing{}, err
	}
	var id string
	var title string
	var lenderId string
	var streetAddress string
	var latitude float32
	var longitude float32
	var feeAmount int64
	var feeCurrency string
	var feeType string
	var imageUrls []string

	for rows.Next() {
		var uid string
		var imageUrl string
		err = rows.Scan(&id, &uid, &title, &lenderId, &streetAddress, &latitude, &longitude, &imageUrl, &feeAmount, &feeCurrency, &feeType)
		imageUrls = append(imageUrls, imageUrl)
	}
	if err != nil {
		return domain.Listing{}, err
	}
	listing, err := domain.ListingRaw{
		Uid:           uid,
		Title:         title,
		LenderId:      lenderId,
		StreetAddress: streetAddress,
		Latitude:      latitude,
		Longitude:     longitude,
		ImageUrls:     imageUrls,
		Fee: domain.FeeRaw{
			Amount: domain.AmountRaw{
				Value:    feeAmount,
				Currency: feeCurrency,
			},
			Type: feeType,
		},
	}.ToEntity()
	if err != nil {
		log.Fatalf(err.Error())
		return domain.Listing{}, nil
	}
	return listing, nil
}

func (r *ListingRepository) FindManyByLatLng(latitude float32, longitude float32, distance int32) ([]domain.Listing, error) {
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
			x GROUP BY x.id, x.uid, x.title, x.lender_id, x.street_address, x.latitude, x.longitude, x.url, x.amount, x.currency, x.type, x.distance 
			HAVING x.distance < $3 ORDER BY x.distance LIMIT 10
		`,
		latitude, longitude, distance,
	)
	if err != nil {
		return []domain.Listing{}, err
	}
	listingsMap := map[string]domain.Listing{}
	for rows.Next() {
		var id string
		var uid string
		var title string
		var lenderId string
		var streetAddress string
		var latitude float32
		var longitude float32
		var distance float32
		var imageUrl string
		var feeAmount int64
		var feeCurrency string
		var feeType string
		err := rows.Scan(&id, &uid, &title, &lenderId, &streetAddress, &latitude, &longitude, &distance, &imageUrl, &feeAmount, &feeCurrency, &feeType)
		if err != nil {
			return []domain.Listing{}, err
		}
		if entry, ok := listingsMap[uid]; ok {
			entry.ImageUrls = append(listingsMap[uid].ImageUrls, imageUrl)
		} else {
			listing, err := domain.ListingRaw{
				Uid:           uid,
				Title:         title,
				LenderId:      lenderId,
				StreetAddress: streetAddress,
				Latitude:      latitude,
				Longitude:     longitude,
				ImageUrls:     append([]string{}, imageUrl),
				Fee: domain.FeeRaw{
					Amount: domain.AmountRaw{
						Value:    feeAmount,
						Currency: feeCurrency,
					},
					Type: feeType,
				},
			}.ToEntity()
			if err != nil {
				log.Fatalf(err.Error())
				return []domain.Listing{}, err
			}
			listingsMap[uid] = listing
		}
	}
	listings := []domain.Listing{}
	for _, value := range listingsMap {
		listings = append(listings, value)
	}
	return listings, nil
}