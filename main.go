package main

import (
	"github.com/cebilon123/nbp-go/internal/logg/nbp"
	"github.com/cebilon123/nbp-go/pkg/checker"
)

func main() {
	l := nbp.NewDefaultLoggerAggregator()
	ch := checker.NewChecker(l, 10, 5)
	if err := ch.Start(); err != nil {
		panic(err)
	}
}
