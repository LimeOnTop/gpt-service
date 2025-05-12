package app

import (
	"fmt"
	"gpt-service/config"
	"gpt-service/gen/gpt"
	"net"

	"context"
	"gpt-service/internal/adapter/token"
	"gpt-service/internal/controller/grpc/gpt"
	"gpt-service/internal/handler"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"

	"log"
	"os"
)

func Run(cfg *config.Config, devMode bool) {
	logger := log.Default()

	// Устанавливаем переменную окружения
	err := os.Setenv("GPT_AUTHORIZATION_KEY", cfg.Token.AuthorizationKey)
	if err != nil {
		logger.Fatalf("Failed to set env variable: %v", err)
	}

	// Проверяем, что переменная установлена
	_, ok := os.LookupEnv("GPT_AUTHORIZATION_KEY")
	if !ok {
		logger.Fatal("GPT_AUTHORIZATION_KEY not found in environment")
	}

	token, err := token.Auth()
	if err != nil {
		logger.Fatal("Failed to get authentication token with error: ", err)
	}

	gptHandler := handler.New(token)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcLogUnaryInterceptor),
		grpc.StreamInterceptor(grpcLogStreamInterceptor),
	)

	grpcController := grpcgpt.New(gptHandler)
	gpt.RegisterRecommendationServer(grpcServer, grpcController)

	//Слушаем порт grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		logger.Fatalf("Failed to listen on port %d: %v", cfg.GRPC.Port, err)
	}
	logger.Printf("gRPC server listening on port %d", cfg.GRPC.Port)
	// Запускаем gRPC сервер
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve gRPC server: %v", err)
	}
}

// Интерсепторы для логирования
func grpcLogStreamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger := log.Default()
	logger.Printf("gRPC Stream called: %s from %s", info.FullMethod, ss.Context().Value("peer").(*peer.Peer).Addr.String())
	return handler(srv, ss)
}

func grpcLogUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	logger := log.Default()
	peerInfo, _ := peer.FromContext(ctx)
	//if !ok {
	//	logger.Printf("gRPC Unary called: %s from UNKNOWN", info.FullMethod)
	//} else {
	//	//logger.Printf("gRPC Unary called: %s from %s", info.FullMethod, peerInfo.Addr.String())
	//}

	resp, err := handler(ctx, req)

	if err != nil {
		logger.Printf("gRPC Unary response error: %s, method: %s, error: %v", peerInfo.Addr.String(), info.FullMethod, err)
	} else {
		logger.Printf("gRPC Unary response: %s, method: %s, response: %v", peerInfo.Addr.String(), info.FullMethod, resp)
	}

	return resp, err
}
