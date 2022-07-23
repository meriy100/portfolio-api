package repositories

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
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
	//ctx := context.Background()
	//
	//client, err := getFireBaseClient(ctx)
	//if err != nil {
	//	fmt.Printf("Failed getFireBaseClient(ctx): %v\n", err)
	//	return err
	//}

	_, err := pr.client.Collection("portfolio-data-profile").Doc("1").Set(pr.ctx, profile)
	if err != nil {
		fmt.Printf("Failed client.Collection: %v\n", err)
		return err
	}
	return nil
}
