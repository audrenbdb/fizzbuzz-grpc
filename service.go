//package algorithm exposes public endpoints to solve common algorithms
package algorithm

import (
	"context"

	"github.com/audrenbdb/algorithm-microsrv/fizzbuzz"
)

type service struct{}

func NewService() *service {
	return &service{}
}

type Service interface {
	FizzBuzz(ctx context.Context, encodedPuzzleLength int32) (string, error)
}

func (s *service) FizzBuzz(_ context.Context, puzzleLength int32) (string, error) {
	return fizzbuzz.Solve(int(puzzleLength)), nil
}
