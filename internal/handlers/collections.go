package handlers

import (
	"errors"
	collection_model "krate/internal/model"

	"github.com/objectbox/objectbox-go/objectbox"
)

// CollectionHandler handles collection-related operations
type CollectionHandler struct {
	ob *objectbox.ObjectBox
}

// NewCollectionHandler creates a new CollectionHandler
func NewCollectionHandler(ob *objectbox.ObjectBox) *CollectionHandler {
	return &CollectionHandler{ob: ob}
}

// GetCollections returns all collections from the database
func (h *CollectionHandler) GetCollections() ([]collection_model.Collection, error) {
	if h.ob == nil {
		return nil, errors.New("ObjectBox not initialized")
	}

	box := collection_model.BoxForCollection(h.ob)
	collections, err := box.GetAll()
	if err != nil {
		return nil, err
	}

	// Convert []*Collection to []Collection
	result := make([]collection_model.Collection, len(collections))
	for i, c := range collections {
		result[i] = *c
	}

	return result, nil
}

// GetCollection returns a single collection by ID
func (h *CollectionHandler) GetCollection(id uint64) (*collection_model.Collection, error) {
	if h.ob == nil {
		return nil, errors.New("ObjectBox not initialized")
	}

	box := collection_model.BoxForCollection(h.ob)
	return box.Get(id)
}

// CreateCollection creates a new collection
func (h *CollectionHandler) CreateCollection(name, genre, imageUrl string, progress float64) (*collection_model.Collection, error) {
	if h.ob == nil {
		return nil, errors.New("ObjectBox not initialized")
	}

	box := collection_model.BoxForCollection(h.ob)
	collection := &collection_model.Collection{
		Name:     name,
		Genre:    genre,
		ImageUrl: imageUrl,
		Progress: progress,
	}

	id, err := box.Put(collection)
	if err != nil {
		return nil, err
	}

	collection.Id = id
	return collection, nil
}

// UpdateCollection updates an existing collection
func (h *CollectionHandler) UpdateCollection(collection *collection_model.Collection) error {
	if h.ob == nil {
		return errors.New("ObjectBox not initialized")
	}

	box := collection_model.BoxForCollection(h.ob)
	_, err := box.Put(collection)
	return err
}

// DeleteCollection deletes a collection by ID
func (h *CollectionHandler) DeleteCollection(id uint64) error {
	if h.ob == nil {
		return errors.New("ObjectBox not initialized")
	}

	box := collection_model.BoxForCollection(h.ob)
	collection, err := box.Get(id)
	if err != nil {
		return err
	}
	if collection == nil {
		return errors.New("collection not found")
	}

	return box.Remove(collection)
}

