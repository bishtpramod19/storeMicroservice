package db

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/bishtpramod19/microservices/order/order/internal/application/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	DatabaseUrl string
}

func (O *OrderDatabaseTestSuite) SetupSuite() {

	ctx := context.Background()
	port := "3306/tcp"
	// dbUrl := func(port nat.Port) string {
	// 	return fmt.Sprintf("root:s3cr3t@tcp(localhost:%s)/orders?charset=utf8mb4&parseTime=True&loc=Local", port.Port())

	// }

	req := testcontainers.ContainerRequest{
		Image:        "docker.io/mysql:8.0.30",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "s3cr3t",
			"MYSQL_DATABASE":      "orders",
		},
		WaitingFor: wait.ForAll(
			wait.ForLog("port: 3306  MySQL Community Server - GPL"),
			wait.ForExposedPort().WithStartupTimeout(180*time.Second),
			wait.ForListeningPort("3306/tcp").WithStartupTimeout(10*time.Second),
		).WithStartupTimeoutDefault(120 * time.Second).WithDeadline(360 * time.Second),
	}

	mysqlContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		log.Fatal("Failed to start Mysql.", err)
	}

	endPoint, _ := mysqlContainer.Endpoint(ctx, "")
	O.DatabaseUrl = fmt.Sprintf("root:s3cr3t@tcp(%s)/orders?charset=utf8mb4&parseTime=True&loc=Local", endPoint)

}

func (O *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	adapter, err := NewAdapter(O.DatabaseUrl)
	O.Nil(err)
	saveErr := adapter.Save(&domain.Order{})
	O.Nil(saveErr)
}

func (O *OrderDatabaseTestSuite) Test_Should_Get_Order() {
	adapter, _ := NewAdapter(O.DatabaseUrl)
	order := domain.NewOrder(2, []domain.OrderItem{
		{
			ProductCode: "CAM",
			Quantity:    5,
			UnitPrice:   1.32,
		},
	})
	adapter.Save(&order)

	ord, _ := adapter.Get(strconv.Itoa(int(order.Id)))
	O.Equal(int64(2), ord.CustomerId)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDatabaseTestSuite))
}
