//package algorithm exposes public endpoints to solve common algorithms
package algorithm

import (
	"context"

	"github.com/audrenbdb/algorithm-microsrv/pb"
	gt "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	fizzBuzz gt.Handler
	history  gt.Handler
}

func NewGRPCServer(endpoints Endpoints) pb.SolveServer {

	return &grpcServer{
		fizzBuzz: gt.NewServer(endpoints.FizzBuzz,
			decodeFizzBuzzRequest,
			encodeFizzBuzzResponse),
		history: gt.NewServer(
			endpoints.History,
			decodeHistoryRequest,
			encodeHistoryResponse,
		),
	}
}

func (s *grpcServer) FizzBuzz(ctx context.Context, req *pb.FizzBuzzRequest) (*pb.FizzBuzzResponse, error) {
	_, resp, err := s.fizzBuzz.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.FizzBuzzResponse), nil
}

func (s *grpcServer) History(ctx context.Context, req *pb.HistoryRequest) (*pb.HistoryResponse, error) {
	_, resp, err := s.history.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.HistoryResponse), nil
}

func decodeHistoryRequest(_ context.Context, request interface{}) (interface{}, error) {
	return HistoryRequest{}, nil
}

func encodeHistoryResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(HistoryResponse)
	resolutions := []*pb.ResolutionResponse{}
	for _, resolution := range resp.Resolutions {
		resolutions = append(resolutions, &pb.ResolutionResponse{
			DateResolved: resolution.dateResolved.Format("Mon 2 Jan 15:04:05"),
			PuzzleName:   resolution.puzzleName,
			Solution:     resolution.solution,
		})
	}
	return &pb.HistoryResponse{Resolutions: resolutions}, nil
}

func decodeFizzBuzzRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.FizzBuzzRequest)
	return FizzBuzzRequest{
		EncodedPuzzleLength: req.PuzzleLength,
	}, nil
}

func encodeFizzBuzzResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(FizzBuzzResponse).Resolution
	return &pb.FizzBuzzResponse{Resolution: &pb.ResolutionResponse{
		DateResolved: resp.dateResolved.Format("Mon 2 Jan 15:04:05"),
		PuzzleName:   resp.puzzleName,
		Solution:     resp.solution,
	}}, nil
}
