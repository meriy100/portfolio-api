package main

import (
	"context"
	"fmt"
	"github.com/meriy100/portfolio-api/adapters"
	"github.com/meriy100/portfolio-api/interfaces/controllers"
	controllerCli "github.com/meriy100/portfolio-api/interfaces/controllers/cli"
	"github.com/meriy100/portfolio-api/interfaces/presenters"
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

	profileCli := controllers.NewProfileCli(
		repositories.NewPostRepository(),
		repositories.NewProfileRepository(ctx, firestoreClient),
		usecase.NewProfileInteractor,
		presenters.NewProfileCliPresenter,
	)

	historyCli := controllerCli.NewHistoryController(
		repositories.NewHistoryRepository(ctx, firestoreClient),
		repositories.NewPostRepository(),
		usecase.NewHistoryInteractor,
		presenterCli.NewHistoryPresenter,
	)

	profileCli.UpdateProfile()
	profileCli.ShowProfile()

	historyCli.UpdateHistories()
	historyCli.IndexHistories()
}
