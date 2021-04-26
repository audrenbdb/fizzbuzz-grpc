//package algorithm expose public endpoints to solve common algorithms
package algorithm

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	FizzBuzz endpoint.Endpoint
}

type FizzBuzzRequest struct {
	EncodedPuzzleLength int32
}

type FizzBuzzResponse struct {
	DecodedPuzzle string
}

func CreateEndpoints(s *service) Endpoints {
	return Endpoints{
		FizzBuzz: createFizzBuzzEndpoint(s),
	}
}

func createFizzBuzzEndpoint(s *service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FizzBuzzRequest)
		res, _ := s.FizzBuzz(ctx, req.EncodedPuzzleLength)
		return FizzBuzzResponse{DecodedPuzzle: res}, nil
	}
}
