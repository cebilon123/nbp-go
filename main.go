package main

import (
	"github.com/cebilon123/nbp-go/internal/logg/nbp"
	"github.com/cebilon123/nbp-go/pkg/checker"
	"github.com/mattn/go-tty"
	"log"
)

const host = "http://api.nbp.pl/api/exchangerates/rates/a/eur/last/100/?format=json"

func main() {
	l := nbp.NewDefaultLoggerAggregator()
	ch := checker.NewChecker(l, 10, 5, host)
	t, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func(t *tty.TTY) {
		err := t.Close()
		if err != nil {
			log.Println(err)
		}
	}(t)

	go func() {
		for {
			r, err := t.ReadRune()
			if err != nil {
				log.Fatal(err)
			}

			// close checker by clicking any button
			if r != 0 {
				ch.Close()
			}
		}
	}()

	go ch.Start()

	ch.Wait()
}
