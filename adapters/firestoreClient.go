package adapters

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

func InitialFireStoreClient(ctx context.Context) (*firestore.Client, error) {
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "portfolio-357112"})
	if err != nil {
		fmt.Printf("Failed firebase.NewApp(ctx, nil, opt): %v\n", err)
		return nil, err
	}

	return app.Firestore(ctx)
}
