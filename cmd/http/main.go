package main

import (
	"github.com/meriy100/portfolio-api/router"
	"net/http"
)

func main() {
	http.HandleFunc("/profile", router.Profile)
	http.HandleFunc("/histories", router.Histories)
	http.HandleFunc("/skills", router.Skills)
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}

}
