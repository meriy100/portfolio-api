package p

import (
	"fmt"
	"github.com/meriy100/portfolio-api/adapters"
	"github.com/meriy100/portfolio-api/interfaces/controllers"
	"github.com/meriy100/portfolio-api/interfaces/presenters"
	"github.com/meriy100/portfolio-api/interfaces/repositories"
	"github.com/meriy100/portfolio-api/usecase"
	"net/http"
)

func UpdatePortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	ctx := r.Context()

	firestoreClient, err := adapters.InitialFireStoreClient(ctx, "serverless_function_source_code/serviceAccountKey.json")

	if err != nil {
		http.Error(w, fmt.Sprintf("Error initial firestore client : %v\n", err.Error()), http.StatusInternalServerError)
		return
	}

	profileController := controllers.NewProfileController(
		repositories.NewPostRepository(),
		repositories.NewProfileRepository(ctx, firestoreClient),
		usecase.NewProfileInteractor,
		presenters.NewProfileHttpPresenter,
	)

	fmt.Printf("query parameter : %s\n", r.FormValue("t"))

	switch r.Method {
	case http.MethodGet:
		switch r.FormValue("t") {
		case "portfolio":
			profileController.ShowProfile(w, r)
		case "histories":
		default:
			profileController.ShowProfile(w, r)
		}

	case http.MethodPost:
		profileController.UpdateProfile(w, r)
	}
}
