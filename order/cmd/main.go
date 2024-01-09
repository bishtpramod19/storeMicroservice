package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"os"

	"github.com/bishtpramod19/microservices/order/order/config"
	"github.com/bishtpramod19/microservices/order/order/internal/adapters/db"
	"github.com/bishtpramod19/microservices/order/order/internal/adapters/grpc"
	"github.com/bishtpramod19/microservices/order/order/internal/adapters/payment"
	"github.com/bishtpramod19/microservices/order/order/internal/application/core/api"
	"google.golang.org/grpc/credentials"
)

func main() {
	dbadapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("failed to connect to database url : %v", err)
	}

	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		log.Fatalf("failed to initialize payment stub. Error: %v", err)
	}
	application := api.NewApplication(dbadapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()

}

func getTlsCredentials() (credentials.TransportCredentials, error) {
	serverCert, _ := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	certpool := x509.NewCertPool()
	caCert, _ := os.ReadFile("cert/ca-cert.pem")

	if ok := certpool.AppendCertsFromPEM(caCert); !ok {
		return nil, errors.New("failed to append the CA certs")
	}

	return credentials.NewTLS(
		&tls.Config{
			ClientAuth:   tls.RequireAnyClientCert,
			Certificates: []tls.Certificate{serverCert},
			ClientCAs:    certpool,
		}), nil
}
