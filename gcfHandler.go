package p

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/meriy100/portfolio-api/router"
	"net/http"
)

func init() {
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
