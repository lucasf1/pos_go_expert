package usecase

import "github.com/lucasf1/Desafio3/internal/entity"

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(repo entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: repo,
	}
}

func (c *ListOrdersUseCase) Execute() ([]OrderOutputDTO, error) {

	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var outputs []OrderOutputDTO

	for _, o := range orders {

		outputs = append(outputs, OrderOutputDTO{
			ID:         o.ID,
			Price:      o.Price,
			Tax:        o.Tax,
			FinalPrice: o.FinalPrice,
		})
	}

	return outputs, nil
}
