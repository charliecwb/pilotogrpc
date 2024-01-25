package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	listener net.Listener
	server   *grpc.Server
	logger   *zap.Logger
)

func main() {
	log, _ := zap.NewProduction()
	defer log.Sync()

	initListener()

	server = grpc.NewServer()

	// Register gRPC server
	//greeting_v1.RegisterGreetServiceServer(server, &greeting_v1.UnimplementedGreetServiceServer{})

	go signalsListener(server)

	if err := server.Serve(listener); err != nil {
		log.Panic("failed to serve", zap.Error(err))
	}
}

func initListener() {
	var err error
	listener, err = net.Listen("tcp", ":8080")
	if err != nil {
		logger.Panic("failed to listen", zap.Error(err))
	}

	logger.Info("listening on port 8080")
}

func signalsListener(server *grpc.Server) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	_ = <-sigs

	logger.Info("shutting down gRPC server...")
	server.GracefulStop()
}
