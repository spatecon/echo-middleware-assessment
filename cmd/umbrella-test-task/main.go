package main

import (
	"github.com/rs/zerolog/log"

	app "github.com/spatecon/echo-middleware-assessment/internal/app/umbrella-test-task"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
