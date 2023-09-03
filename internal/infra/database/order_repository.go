package database

import (
	"database/sql"
	"fmt"

	"github.com/patyumi/api-orders/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll(page, limit int, sort string) ([]entity.Order, error) {
	var orders []entity.Order
	var query string
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		query = fmt.Sprintf("Select id, price, tax, final_price from orders order by id %s limit %d,%d", sort, page-1, limit)
	} else {
		query = fmt.Sprintf("Select id, price, tax, final_price from orders order by id %s", sort)
	}

	rows, err := r.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o entity.Order
		err = rows.Scan(&o.ID, &o.Price, &o.Tax, &o.FinalPrice)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	return orders, err
}
