package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/cortezzIP/Simbir.Health-API/gen/go/hospital-service"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/config"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/database"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/repository"
	"github.com/cortezzIP/Simbir.Health-API/services/hospital-service/internal/service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	cfg, err := config.NewConfig("../config/config.yaml")
	if err != nil {
		slog.Error("Failed to read config: " + err.Error())
		os.Exit(1)
	}

	database.ConnectToHospitalDB(&cfg.Database)
	defer database.CloseHospitalDB()
	
	lis, err := net.Listen("tcp", cfg.GRPC.Port)
	if err != nil {
		slog.Error("Failed to listen: " + err.Error())
		os.Exit(1)
	}

	grpcServer := grpc.NewServer()
	service := service.NewHospitalServiceServer(repository.NewHospitalRepository())
	
	pb.RegisterHospitalServer(grpcServer, service)
	
	go func() {
		slog.Info("gRPC hospital server started on " + cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("Failed to start gRPC server: " + err.Error())
			os.Exit(1)
		}
	}()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	
	mux := runtime.NewServeMux()
	
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
  	err = pb.RegisterHospitalHandlerFromEndpoint(ctx, mux, cfg.GRPC.Port, opts)
	if err != nil {
		slog.Error("Failed to register hospital handler: " + err.Error())
		os.Exit(1)
	}

	slog.Info("HTTP hospital server started on " + cfg.GRPCGateway.Port)
	if err = http.ListenAndServe(cfg.GRPCGateway.Port, mux); err != nil {
		slog.Error("Failed to start http server: " + err.Error())
		os.Exit(1)
	}
}