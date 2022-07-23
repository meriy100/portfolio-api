package presenters

import (
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type ProfileCliPresenter struct {
}

func NewProfileCliPresenter() ports.ProfileOutputPort {
	return &ProfileCliPresenter{}
}

func (pp *ProfileCliPresenter) OutputFetchPostError(err error) error {
	fmt.Printf("OutputFetchPostError: %v\n", err)
	return nil
}

func (pp *ProfileCliPresenter) OutputFindProfileError(err error) error {
	fmt.Printf("OutputFindProfileError: %v\n", err)
	return nil
}

func (pp *ProfileCliPresenter) OutputToProfileError(err error) error {
	fmt.Printf("OutputToProfileError: %v\n", err)
	return nil
}

func (pp *ProfileCliPresenter) OutputProfileSaveError(err error) error {
	fmt.Printf("OutputProfileSaveError: %v\n", err)
	return nil
}

func (pp *ProfileCliPresenter) OutputSuccessUpdate() error {
	fmt.Printf("Success Update Profile!")
	return nil
}

func (pp *ProfileCliPresenter) OutputProfile(profile *entities.Profile) error {
	fmt.Printf("Profile : %v\n", profile)
	return nil
}
