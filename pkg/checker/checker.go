package checker

import "github.com/cebilon123/nbp-go/internal/logg"

// checker is used to check
// state of NBP currencies and
// logs it to given logger targets
type checker struct {
	logger logg.Logger
	checkerConfig
}

// checkerConfig represents configuration
// for checker.
type checkerConfig struct {
	checksAmount  int // checksAmount - how many checks are going to be made in given time span
	secondsAmount int // secondsAmount - time span
}

func NewChecker(logger logg.Logger, checksAmount int, secondsAmount int) *checker {
	return &checker{logger: logger}
}

// Start starts checker which will
// do checks based on passed config.
func (c *checker) Start() error {
	return nil
}
