package main

import (
	"context"
	"github.com/cebilon123/nbp-go/internal/logg"
	"github.com/cebilon123/nbp-go/internal/logg/console"
	"github.com/cebilon123/nbp-go/internal/logg/file"
	"github.com/cebilon123/nbp-go/internal/logg/nbp"
	"github.com/cebilon123/nbp-go/pkg/checker"
	"github.com/mattn/go-tty"
	"log"
)

const host = "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	l := nbp.NewLoggerAggregator([]logg.Logger{console.NewConsoleLogger(), file.NewFileLogger(ctx)})
	ch := checker.NewChecker(l, 10, 5, host, ctx, cancel)
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	// close tty
	defer func(t *tty.TTY) {
		err := t.Close()
		if err != nil {
			log.Println(err)
		}
	}(t)

	// listen for key pushes
	go func() {
		for {
			r, err := t.ReadRune()
			if err != nil {
				log.Fatal(err)
			}

			// close checker by clicking any key
			if r != 0 {
				ch.Close()
			}
		}
	}()

	// start checker in another goroutine
	go ch.Start()

	// wait for checker to be done:
	// 1. there was a panic
	// 2. it was closed by user
	ch.Wait()
}
