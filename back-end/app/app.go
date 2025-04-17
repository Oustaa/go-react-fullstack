package app

import "log"

type Application struct {
	Logger log.Logger
}

func NewApplication() *Application {
	app := Application{}

	return &app
}
