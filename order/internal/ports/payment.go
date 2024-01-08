package ports

import "github.com/bishtpramod19/microservices/order/order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
