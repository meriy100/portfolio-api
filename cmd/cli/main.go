package main

import (
	"context"
	"fmt"

	"github.com/meriy100/portfolio-api/adapters"
	controllerCli "github.com/meriy100/portfolio-api/interfaces/controllers/cli"
	presenterCli "github.com/meriy100/portfolio-api/interfaces/presenters/cli"
	"github.com/meriy100/portfolio-api/interfaces/repositories"
	"github.com/meriy100/portfolio-api/usecase"
)

func main() {
	ctx := context.Background()
	firestoreClient, err := adapters.InitialFireStoreClient(ctx)

	if err != nil {
		panic(fmt.Errorf("error initial firestore client : %w", err))
	}

	profileController := controllerCli.NewProfileCli(
		repositories.NewPostRepository(),
		repositories.NewProfileRepository(ctx, firestoreClient),
		repositories.NewContentDeliveryRepository(),
		usecase.NewProfileInteractor,
		presenterCli.NewProfilePresenter,
	)

	historyController := controllerCli.NewHistoryController(
		repositories.NewHistoryRepository(ctx, firestoreClient),
		repositories.NewPostRepository(),
		usecase.NewHistoryInteractor,
		presenterCli.NewHistoryPresenter,
	)

	skillController := controllerCli.NewSkillController(
		repositories.NewSkillRepository(ctx, firestoreClient),
		repositories.NewPostRepository(),
		usecase.NewSkillInteractor,
		presenterCli.NewSkillPresenter,
	)

	profileController.UpdateProfile()
	profileController.ShowProfile()

	historyController.UpdateHistories()
	historyController.IndexHistories()

	skillController.UpdateSkills()
	skillController.IndexSkills()
}
