package p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Post struct {
	BodyMd string `json:"body_md"`
}

func fetchPost() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.esa.io/v1/teams/meriy100/posts/253", nil)
	if err != nil {
		return []byte{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ESA_ACCESS_TOKEN")))
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return byteArray, nil
}

func FetchPortfolio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "3600")

	body, err := fetchPost()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//var d struct {
	//	Message string `json:"message"`
	//}

	//if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
	//	switch err {
	//	case io.EOF:
	//		fmt.Fprint(w, "Hello World!")
	//		return
	//	default:
	//		log.Printf("json.NewDecoder: %v", err)
	//		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	//		return
	//	}
	//}

	fmt.Fprint(w, string(body))
	//if d.Message == "" {
	//	fmt.Fprint(w, "Hello World!")
	//	return
	//}
	//fmt.Fprint(w, html.EscapeString(d.Message))
}

func main() {
	body, err := fetchPost()
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(post)
}
