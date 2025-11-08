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
	collectionHandler *handlers.CollectionHandler
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	objectBox, err := objectbox.NewBuilder().
		Model(collection_model.ObjectBoxModel()).
		Build()

	if err != nil {
		// A real app should handle this more gracefully
		log.Fatalf("Failed to build ObjectBox: %s", err)
	}

	a.ob = objectBox
	a.collectionHandler = handlers.NewCollectionHandler(objectBox)
	log.Println("ObjectBox database initialized successfully.")
}

func (a *App) shutdown(ctx context.Context) {
	if a.ob != nil {
		a.ob.Close()
		log.Println("ObjectBox database closed.")
	}
}
