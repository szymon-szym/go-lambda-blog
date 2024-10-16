package items

import (
	"database/sql"
	"lambda_handlers/api"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestItemService_GetItem_Success(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.NoError(t, err)

	defer db.Close()

	itemService := InitializeItemsService(db)

	expectedItem := api.Item{
		Id:    1,
		Name:  "test name",
		Price: 11.11,
	}

	mock.ExpectQuery("SELECT \\* FROM items WHERE id = \\$1").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "price"}).AddRow(expectedItem.Id, expectedItem.Name, expectedItem.Price))

	item, err := itemService.GetItem(1)

	assert.NoError(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, &expectedItem, item)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestItemService_GetItem_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()

	assert.NoError(t, err)

	defer db.Close()

	itemService := InitializeItemsService(db)

	mock.ExpectQuery("SELECT \\* FROM items WHERE id = \\$1").WithArgs(42).WillReturnError(sql.ErrNoRows)

	item, err := itemService.GetItem(42)

	assert.Error(t, err)
	assert.Nil(t, item)
	assert.Equal(t, "Item not found", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())

}
