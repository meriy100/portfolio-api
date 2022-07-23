package adapters

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
)

func InitialFireStoreClient(ctx context.Context, credentialFilePath string) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(credentialFilePath)
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "portfolio-357112"}, opt)
	if err != nil {
		fmt.Printf("Failed firebase.NewApp(ctx, nil, opt): %v\n", err)
		return nil, err
	}

	return app.Firestore(ctx)
}
