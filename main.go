package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/jicol95/currency-service/protos/currency"
	"github.com/jicol95/currency-service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := hclog.Default()
	server := grpc.NewServer()
	cs := service.NewCurrency(logger)

	protos.RegisterCurrencyServer(server, cs)
	reflection.Register(server)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		logger.Error("Unable to listen on port 9092", "error", err)
		os.Exit(1)
	}

	server.Serve(l)
}
