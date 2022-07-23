package ports

import (
	"github.com/meriy100/portfolio-api/entities"
)

type ProfileRepository interface {
	Save(profile *entities.Profile) error
}

type ProfileInputPort interface {
	UpdateProfile() error
}

type ProfileOutputPort interface {
	OutputFetchPostError(error) error
	OutputToProfileError(error) error
	OutputProfileSaveError(error) error

	OutputSuccessUpdate() error
}
