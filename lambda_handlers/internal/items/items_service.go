package items

import (
	"database/sql"
	"errors"
	api "lambda_handlers/api"
	"log"
)

type ItemsService struct {
	conn *sql.DB
}

func InitializeItemsService(conn *sql.DB) *ItemsService {
	return &ItemsService{
		conn: conn,
	}
}

func (svc ItemsService) GetItem(id int) (*api.Item, error) {

	log.Printf("Getting item id %v", id)

	query := `SELECT * FROM items WHERE id = $1`

	var item api.Item

	err := svc.conn.QueryRow(query, id).Scan(&item.Id, &item.Name, &item.Price)

	log.Printf("Got item id %v", id)

	if err != nil {
		log.Printf("Error getting item %v: %v", id, err)
		if err == sql.ErrNoRows {
			return nil, errors.New("Item not found")
		}
		return nil, err
	}

	return &item, nil

}
