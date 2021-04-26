//package algorithm exposes public endpoints to solve common algorithms
package algorithm

import (
	"context"

	"github.com/audrenbdb/algorithm-microsrv/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	add gt.Handler
}

func NewGRPCServer(endpoints Endpoints) pb.SolveServer {
	return &grpcServer{
		add: gt.NewServer(
			endpoints.FizzBuzz,
			decodeFizzBuzzRequest,
			encodeFizzBuzzResponse,
		),
	}
}

func (s *grpcServer) FizzBuzz(ctx context.Context, req *pb.FizzBuzzRequest) (*pb.FizzBuzzResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.FizzBuzzResponse), nil
}

func decodeFizzBuzzRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FizzBuzzRequest)
	return FizzBuzzRequest{
		EncodedPuzzleLength: req.PuzzleLength,
	}, nil
}

func encodeFizzBuzzResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(FizzBuzzResponse)
	return &pb.FizzBuzzResponse{Solution: resp.DecodedPuzzle}, nil
}
