package repository

import (
	"database/sql"

	domain "github.com/kokiebisu/rental-storage/service-user/internal/core/domain/user"
	_ "github.com/lib/pq"
)

type ItemRepository struct {
	db *sql.DB
}

func NewItemRepository(db *sql.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

func (r *ItemRepository) Setup() error {
	// need ton convert this to transaction
	_, err := r.db.Exec(`
		CREATE TABLE IF NOT EXISTS item (
			id SERIAL NOT NULL PRIMARY KEY, 
			uid VARCHAR(64) NOT NULL,
			name VARCHAR(20) NOT NULL, 
			owner_id VARCHAR(30) NOT NULL, 
			listing_id VARCHAR(30) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
			updated_at TIMESTAMP,
			CONSTRAINT fk_user
			FOREIGN KEY(owner_id) 
				REFERENCES user(id)
					ON DELETE CASCADE
	  	)
	`)
	if err != nil {
		// rollback
		return err
	}
	_, err = r.db.Exec(`
		CREATE TABLE IF NOT EXISTS item_images (
			id SERIAL NOT NULL PRIMARY KEY,
			url VARCHAR(150) NOT NULL,
			item_id INT NOT NULL,
			CONSTRAINT fk_item
			FOREIGN KEY(item_id) 
				REFERENCES item(id)
					ON DELETE CASCADE
		)
	`)
	if err != nil {
		// rollback
		return err
	}
	return nil
}

func (r *ItemRepository) Save(item domain.Item) error {
	// begin transaction
	result, err := r.db.Exec(
		`INSERT INTO item (name, owner_id, listing_id, created_at) VALUES($1, $2, $3, $4)`,
		item.Name, item.OwnerId, item.ListingId, item.CreatedAt,
	)
	if err != nil {
		// rollback
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	for _, url := range item.ImageUrls {
		r.db.Exec(
			`INSERT INTO item_images (item_id, image_url) VALUES($1, $2)`,
			id, url,
		)
	}
	if err != nil {
		// rollback
		return err
	}
	return nil
}

func (r *ItemRepository) Delete(uid string) error {
	_, err := r.db.Exec(
		`DELETE FROM item WHERE id = $1`,
		uid,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *ItemRepository) FindOneById(id int64) (domain.Item, error) {
	var uid string
	var name string
	var ownerId string
	var listingId string
	var createdAt string
	var updatedAt string
	row := r.db.QueryRow(
		`SELECT * FROM item WHERE id = $1`,
		uid,
	)
	row.Scan(&id, &uid, &name, &ownerId, &listingId, &createdAt, &updatedAt)
	rows, err := r.db.Query(
		`SELECT * FROM item_images WHERE item_id = $1`,
		id,
	)
	defer rows.Close()
	if err != nil {
		return domain.Item{}, err
	}
	var imageUrls []string
	for rows.Next() {
		var imageUrl string
		err := rows.Scan(&imageUrl)
		if err != nil {
			return domain.Item{}, err
		}
		imageUrls = append(imageUrls, imageUrl)
	}

	item := domain.ItemRaw{
		Uid:       uid,
		Name:      name,
		ImageUrls: imageUrls,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		OwnerId:   ownerId,
		ListingId: listingId,
	}
	return item.ToEntity(), nil

}
