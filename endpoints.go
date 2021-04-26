//package algorithm expose public endpoints to solve common algorithms
package algorithm

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FizzBuzz endpoint.Endpoint
	History  endpoint.Endpoint
}

type FizzBuzzRequest struct {
	EncodedPuzzleLength int32
}
type FizzBuzzResponse struct {
	Resolution resolution
}

type HistoryRequest struct{}
type HistoryResponse struct {
	Resolutions []resolution
}

func CreateEndpoints(s *service) Endpoints {
	return Endpoints{
		FizzBuzz: createFizzBuzzEndpoint(s),
		History:  createHistoryEndpoint(s),
	}
}

func createFizzBuzzEndpoint(s *service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FizzBuzzRequest)
		res, _ := s.FizzBuzz(ctx, req.EncodedPuzzleLength)
		return FizzBuzzResponse{Resolution: res}, nil
	}
}

func createHistoryEndpoint(s *service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, _ := s.History(ctx)
		return HistoryResponse{Resolutions: res}, nil
	}
}
