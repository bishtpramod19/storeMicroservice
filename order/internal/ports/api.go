package ports

import "github.com/bishtpramod19/storeMicroservice/order/internal/application/core/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
