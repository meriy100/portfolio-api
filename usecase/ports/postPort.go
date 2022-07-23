package ports

import "github.com/meriy100/portfolio-api/entities"

type PostRepository interface {
	FetchPost(postId int) (*entities.Post, error)
}
