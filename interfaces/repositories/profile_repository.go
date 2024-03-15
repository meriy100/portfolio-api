package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/meriy100/portfolio-api/entities"
)

type ProfileRepository struct {
	ctx    context.Context
	client *firestore.Client
}

func NewProfileRepository(ctx context.Context, client *firestore.Client) *ProfileRepository {
	return &ProfileRepository{ctx, client}
}

func (pr *ProfileRepository) Save(profile *entities.Profile) error {
	_, err := pr.client.Collection("portfolio-data-profile").Doc("1").Set(pr.ctx, profile)
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProfileRepository) Find() (*entities.Profile, error) {
	var profile entities.Profile
	dsnap, err := pr.client.Collection("portfolio-data-profile").Doc("1").Get(pr.ctx)
	if err != nil {
		return &profile, err
	}

	err = dsnap.DataTo(&profile)
	return &profile, err
}
