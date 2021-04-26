//package algorithm exposes public endpoints to solve common algorithms
package algorithm

import (
	"context"

	"github.com/audrenbdb/algorithm-microsrv/fizzbuzz"
)

type service struct {
	repo repo
}

func NewService(repo repo) *service {
	return &service{repo: repo}
}

type Service interface {
	FizzBuzz(ctx context.Context, encodedPuzzleLength int32) (resolution, error)
	History(ctx context.Context) ([]resolution, error)
}

func (s *service) FizzBuzz(_ context.Context, puzzleLength int32) (resolution, error) {
	decodedFizzBuzz := fizzbuzz.Solve(int(puzzleLength))
	return s.repo.SaveResolution("FizzBuzz", decodedFizzBuzz), nil
}

func (s *service) History(_ context.Context) ([]resolution, error) {
	return s.repo.GetSolvedPuzzleHistory(), nil
}
