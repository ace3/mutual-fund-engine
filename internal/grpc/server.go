package grpc

import (
	"fmt"
	"log"
	"net"

	pb "github.com/ace3/mutual-fund-engine/pkg/pb/nobi"

	"github.com/ace3/mutual-fund-engine/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

func StartServer(cfg *config.Config, db *gorm.DB) {
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterNobiInvestmentServiceServer(grpcServer, NewNobiInvestmentHandler(db))
	reflection.Register(grpcServer)
	fmt.Printf("gRPC server listening on port %s\n", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
