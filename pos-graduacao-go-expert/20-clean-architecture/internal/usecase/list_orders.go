package usecase

import (
	"github.com/breinoso2006/fullcycle/pos-graduacao-go-expert/20-clean-architecture/internal/entity"
)

type OrdersOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (OrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return OrdersOutputDTO{}, err
	}

	var output []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		output = append(output, dto)
	}
	return OrdersOutputDTO{Orders: output}, nil
}
