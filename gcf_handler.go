package p

import (
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/meriy100/portfolio-api/router"
)

func init() {
	functions.HTTP("profile", Profile)
	functions.HTTP("histories", Histories)
	functions.HTTP("skills", Skills)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	router.Profile(w, r)
}

func Histories(w http.ResponseWriter, r *http.Request) {
	router.Histories(w, r)
}

func Skills(w http.ResponseWriter, r *http.Request) {
	router.Skills(w, r)
}
