package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/meriy100/portfolio-api/entities"
)

type HistoryRepository struct {
	ctx    context.Context
	client *firestore.Client
}

func NewHistoryRepository(ctx context.Context, client *firestore.Client) *HistoryRepository {
	return &HistoryRepository{ctx, client}
}

func (pr *HistoryRepository) Save(history *entities.History) error {
	_, err := pr.client.Collection("portfolio-data-histories").Doc(history.Organization).Set(pr.ctx, history)
	if err != nil {
		return err
	}
	return nil
}
