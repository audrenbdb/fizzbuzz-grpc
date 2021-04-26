package algorithm

import "time"

type resolution struct {
	dateResolved time.Time
	puzzleName   string
	solution     string
}

type (
	resolutionSaver interface {
		SaveResolution(puzzleName, solution string) resolution
	}
	historyGetter interface {
		GetSolvedPuzzleHistory() []resolution
	}
)

type repo interface {
	resolutionSaver
	historyGetter
}

type inMemRepo struct {
	history []resolution
}

func NewInMemRepo() *inMemRepo {
	return &inMemRepo{
		history: []resolution{},
	}
}

func (r *inMemRepo) SaveResolution(puzzleName, solution string) resolution {
	resolution := resolution{
		dateResolved: time.Now(),
		puzzleName:   puzzleName,
		solution:     solution,
	}
	r.history = append(r.history, resolution)
	return resolution
}

func (r *inMemRepo) GetSolvedPuzzleHistory() []resolution {
	return r.history
}
