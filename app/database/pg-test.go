package database

import (
	"fmt"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Order struct {
	tableName  struct{} `pg:"order_table"`
	ID         string   `pg:"id,pk"`
	Item       string   `pg:"item"`
	User       string   `pg:"user"`
	Quantity   int      `pg:"quantity"`
	City       string   `pg:"city"`
	Department string   `pg:"department"`
	Price      int      `pg:"price"`
}

type Report struct {
	tableName struct{}  `pg:"report_table"`
	ID        string    `pg:"id,pk"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
}

func (r Report) String() string {
	return fmt.Sprintf("Report<%s %s %s>", r.ID, r.CreatedAt, r.UpdatedAt)
}

func (o Order) String() string {
	return fmt.Sprintf("Order<%s %s %s>", o.ID, o.Item, o.User)
}

func Basic() {
	opt, err := pg.ParseURL("postgres://postgres:root@localhost:5432/reports?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	err = createSchema(db)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Schema creato")
	}

}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*Order)(nil),
		(*Report)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
