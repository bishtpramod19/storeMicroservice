package api

import (
	"github.com/bishtpramod19/storeMicroservice/order/internal/application/core/domain"
	"github.com/bishtpramod19/storeMicroservice/order/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{
		db: db,
	}

}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}
	return order, nil

}
