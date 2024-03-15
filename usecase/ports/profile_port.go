package ports

import (
	"github.com/meriy100/portfolio-api/entities"
)

type ProfileRepository interface {
	Save(profile *entities.Profile) error
	Find() (*entities.Profile, error)
}

type ProfileInputPort interface {
	ShowProfile() error
	UpdateProfile() error
}

type ProfileOutputPort interface {
	OutputFetchPostError(error) error
	OutputToProfileError(error) error
	OutputProfileSaveError(error) error
	OutputFindProfileError(error) error
	OutputDeployError(error) error

	OutputSuccessUpdate() error

	OutputProfile(*entities.Profile) error
}
