package p

import (
	"encoding/json"
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
	"net/http"
)

func FetchPortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	body, err := entities.FetchPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var post entities.Post
	if err := json.Unmarshal(body, &post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := entities.SaveItem(post.BodyMd); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "ok")
}
