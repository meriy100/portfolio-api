package main

import (
	"context"
	"fmt"
	"github.com/meriy100/portfolio-api/adapters"
	"github.com/meriy100/portfolio-api/interfaces/controllers"
	"github.com/meriy100/portfolio-api/interfaces/presenters"
	"github.com/meriy100/portfolio-api/interfaces/repositories"
	"github.com/meriy100/portfolio-api/usecase"
	"os"
)

func main() {
	//body, err := entities.FetchPost()
	//if err != nil {
	//	fmt.Printf("err: %v\n", err.Error())
	//	return
	//}
	//
	//var post entities.Post
	//if err := json.Unmarshal(body, &post); err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	//fmt.Println(post)
	//if err := entities.SaveItem(post.BodyMd); err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	ctx := context.Background()
	//firestoreClient, err := adapters.InitialFireStoreClient(ctx, "serverless_function_source_code/serviceAccountKey.json")
	firestoreClient, err := adapters.InitialFireStoreClient(ctx, "serviceAccountKey.json")

	if err != nil {
		fmt.Printf("Error initial firestore client : %v\n", err)
		os.Exit(2)
	}

	profileCli := controllers.NewProfileCli(
		repositories.NewPostRepository(),
		repositories.NewProfileRepository(ctx, firestoreClient),
		usecase.NewProfileInteractor,
		presenters.NewProfilePresenter,
	)

	profileCli.UpdateProfile()
}
