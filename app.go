package main

import (
	"context"
	"krate/internal/handlers"
	collection_model "krate/internal/model"
	"log"

	"github.com/objectbox/objectbox-go/objectbox"
)

type App struct {
	ctx               context.Context
	ob                *objectbox.ObjectBox
	CollectionHandler *handlers.CollectionHandler
}

func NewApp() (*App, error) {
	app := &App{}

	// Initialize ObjectBox before binding
	objectBox, err := objectbox.NewBuilder().
		Model(collection_model.ObjectBoxModel()).
		Build()

	if err != nil {
		return nil, err
	}

	app.ob = objectBox
	app.CollectionHandler = handlers.NewCollectionHandler(objectBox)
	log.Println("ObjectBox database initialized successfully.")

	return app, nil
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// ObjectBox is already initialized in NewApp
}

func (a *App) shutdown(ctx context.Context) {
	if a.ob != nil {
		a.ob.Close()
		log.Println("ObjectBox database closed.")
	}
}
