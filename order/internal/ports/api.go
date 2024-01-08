package ports

import "github.com/bishtpramod19/microservices/order/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
