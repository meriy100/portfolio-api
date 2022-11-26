package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
	"io"
	"net/http"
	"os"
)

type PostRepository struct {
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

func (pr *PostRepository) FetchPost(postId int) (*entities.Post, error) {
	var post entities.Post
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.esa.io/v1/teams/meriy100/posts/%d", postId), nil)
	if err != nil {
		return &post, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("ESA_ACCESS_TOKEN")))
	resp, err := client.Do(req)
	if err != nil {
		return &post, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("bad response status code %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return &post, err
	}

	if err := json.Unmarshal(byteArray, &post); err != nil {
		fmt.Printf("err: %v\n", err)
		return &post, err
	}
	return &post, err
}
