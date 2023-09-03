package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	// GetTotal() (int, error)
	FindAll(page, limit int, sort string) ([]Order, error)
}
