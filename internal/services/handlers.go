package services

import (
	"context"
	"fmt"
	pb "pilotoGRPC/pkg/pb/greeting/v1"
)

type GreetService struct {
}

func (s *GreetService) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	return &pb.GreetResponse{
		Resp: fmt.Sprintf("%s, %s", req.Msg.Greeting.String(), req.Msg.Name),
	}, nil

}
