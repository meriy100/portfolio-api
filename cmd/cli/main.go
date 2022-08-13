package main

import (
	"context"
	"fmt"
	"github.com/meriy100/portfolio-api/adapters"
	controllerCli "github.com/meriy100/portfolio-api/interfaces/controllers/cli"
	presenterCli "github.com/meriy100/portfolio-api/interfaces/presenters/cli"
	"github.com/meriy100/portfolio-api/interfaces/repositories"
	"github.com/meriy100/portfolio-api/usecase"
	"os"
)

func main() {
	ctx := context.Background()
	firestoreClient, err := adapters.InitialFireStoreClient(ctx, "serviceAccountKey.json")

	if err != nil {
		fmt.Printf("Error initial firestore client : %v\n", err)
		os.Exit(2)
	}

	profileController := controllerCli.NewProfileCli(
		repositories.NewPostRepository(),
		repositories.NewProfileRepository(ctx, firestoreClient),
		usecase.NewProfileInteractor,
		presenterCli.NewProfilePresenter,
	)

	historyController := controllerCli.NewHistoryController(
		repositories.NewHistoryRepository(ctx, firestoreClient),
		repositories.NewPostRepository(),
		usecase.NewHistoryInteractor,
		presenterCli.NewHistoryPresenter,
	)

	profileController.UpdateProfile()
	profileController.ShowProfile()

	historyController.UpdateHistories()
	historyController.IndexHistories()
}
