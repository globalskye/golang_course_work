package service

import (
	"course_work/pkg/model"
	"course_work/pkg/repository"
)

type OrderService struct {
	repo repository.OrderI
}

func (o OrderService) CreateOrder(card model.UserCard) error {
	newCard := card

	for i, v := range card.Items {
		newCard.Items[i].TotalPrice = v.Quantity * v.Price
	}

	return o.repo.CreateOrder(newCard)
}

func (o OrderService) GetAll() ([]model.Order, error) {

	return o.repo.GetAll()
}

func NewOrderService(repo repository.OrderI) *OrderService {
	return &OrderService{repo: repo}
}
