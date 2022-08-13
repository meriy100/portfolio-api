package router

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/meriy100/portfolio-api/adapters"
	controllerHttp "github.com/meriy100/portfolio-api/interfaces/controllers/http"
	presenterHttp "github.com/meriy100/portfolio-api/interfaces/presenters/http"
	"github.com/meriy100/portfolio-api/interfaces/repositories"
	"github.com/meriy100/portfolio-api/usecase"
	"net/http"
	"os"
)

func setHeader(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTION")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func initialFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	firestoreClient, err := adapters.InitialFireStoreClient(ctx, os.Getenv("SERVICE_ACCOUNT_KEY_PATH"))

	if err != nil {
		return nil, err
	}
	return firestoreClient, nil

}

func Profile(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	ctx := r.Context()
	firestoreClient, err := initialFirestoreClient(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error initial firestore client : %v\n", err.Error()), http.StatusInternalServerError)
		return
	}
	profileController := controllerHttp.NewProfileController(
		repositories.NewPostRepository(),
		repositories.NewProfileRepository(ctx, firestoreClient),
		usecase.NewProfileInteractor,
		presenterHttp.NewProfilePresenter,
	)

	switch r.Method {
	case http.MethodGet:
		profileController.ShowProfile(w, r)
	case http.MethodPost:
		profileController.UpdateProfile(w, r)
	}
}

func Histories(w http.ResponseWriter, r *http.Request) {
	setHeader(w, r)
	ctx := r.Context()
	firestoreClient, err := initialFirestoreClient(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error initial firestore client : %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	historyController := controllerHttp.NewHistoryController(
		repositories.NewHistoryRepository(ctx, firestoreClient),
		repositories.NewPostRepository(),
		usecase.NewHistoryInteractor,
		presenterHttp.NewHistoryPresenter,
	)

	switch r.Method {
	case http.MethodGet:
		historyController.IndexHistories(w, r)
	case http.MethodPost:
		historyController.UpdateHistories(w, r)
	}
}
