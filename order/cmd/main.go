package main

import (
	"log"

	"github.com/bishtpramod19/storeMicroservice/order/config"
	"github.com/bishtpramod19/storeMicroservice/order/internal/adapters/db"
	"github.com/bishtpramod19/storeMicroservice/order/internal/adapters/grpc"
	"github.com/bishtpramod19/storeMicroservice/order/internal/application/core/api"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
