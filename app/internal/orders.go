package mycrud

import (
	"github.com/fracartdev/testing/app/database"
	"github.com/fracartdev/testing/graph/model"
	"github.com/go-pg/pg/v10"
)

// ReadOrder legge un ordine dal db
func ReadOrder(id string) (*model.Order, error) {
	opt, err := pg.ParseURL("postgres://postgres:root@localhost:5432/reports?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	var order = new(database.Order)
	err = db.Model(order).
		Where("id = ?", id).
		Select()

	return &model.Order{
		ID:         order.ID,
		Item:       order.Item,
		User:       order.User,
		Quantity:   order.Quantity,
		City:       order.City,
		Department: order.Department,
		Price:      order.Price,
	}, err
}
