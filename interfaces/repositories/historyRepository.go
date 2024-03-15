package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/meriy100/portfolio-api/entities"
	"google.golang.org/api/iterator"
)

type HistoryRepository struct {
	ctx    context.Context
	client *firestore.Client
}

func NewHistoryRepository(ctx context.Context, client *firestore.Client) *HistoryRepository {
	return &HistoryRepository{ctx, client}
}

func (h *HistoryRepository) Save(history *entities.History) error {
	_, err := h.client.Collection("portfolio-data-histories").Doc(history.Organization).Set(h.ctx, history)
	if err != nil {
		return err
	}
	return nil
}

func (h *HistoryRepository) All() ([]*entities.History, error) {
	var histories []*entities.History
	iter := h.client.Collection("portfolio-data-histories").Documents(h.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return histories, err
		}

		var history entities.History
		err = doc.DataTo(&history)
		if err != nil {
			return histories, err
		}
		histories = append(histories, &history)
	}

	return histories, nil
}
