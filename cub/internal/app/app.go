package app

import (
	"sync"
)

// @PublicValueInstance
type App struct {
	Primitives
	Dependencies
	Adapters
	Usecases
	Actions
	Handlers
	Ports
}

func (app *App) Run() error {
	app.Logger.Info("app startup initiated")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		<-app.Interrupter.Context().Done()

		defer wg.Done()
	}()

	wg.Wait()

	return nil
}
