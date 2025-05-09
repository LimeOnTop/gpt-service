package server

import (
	"context"

	pb "gpt-service/gen/gpt"
	"gpt-service/internal/handler"
)

var _ pb.RecommendationServer = (*GPTServer)(nil)

type GPTServer struct {
	pb.UnimplementedRecommendationServer
	gptHandle handler.Handler
}

func(g *GPTServer) GetGPTRecommendation(ctx context.Context, req *pb.UserRequest) (*pb.GPTResponse, error) {
	message, err := g.gptHandle.GetGPTRecommendation(req.Products, req.Preference)
	if err != nil {
		return nil, err
	}

	imageData, imageFormat, err := g.gptHandle.GetGPTImage(message)
	if err != nil {
		return nil, err
	}
	return &pb.GPTResponse{
		Message: message,
		ImageData: imageData,
		ImageFormat: imageFormat,
	}, nil
}
