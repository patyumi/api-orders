package usecase

import (
	"github.com/patyumi/api-orders/internal/entity"
)

type OrderListInputDTO struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Sort  string `json:"sort"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) List(input OrderListInputDTO) ([]entity.Order, error) {
	output, err := c.OrderRepository.FindAll(input.Page, input.Limit, input.Sort)
	if err != nil {
		return nil, err
	}
	return output, nil
}
