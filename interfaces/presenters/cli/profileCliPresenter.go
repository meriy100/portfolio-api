package cli

import (
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type ProfilePresenter struct {
}

func NewProfilePresenter() ports.ProfileOutputPort {
	return &ProfilePresenter{}
}

func (pp *ProfilePresenter) OutputFetchPostError(err error) error {
	fmt.Printf("OutputFetchPostError: %v\n", err)
	return nil
}

func (pp *ProfilePresenter) OutputFindProfileError(err error) error {
	fmt.Printf("OutputFindProfileError: %v\n", err)
	return nil
}

func (pp *ProfilePresenter) OutputToProfileError(err error) error {
	fmt.Printf("OutputToProfileError: %v\n", err)
	return nil
}

func (pp *ProfilePresenter) OutputProfileSaveError(err error) error {
	fmt.Printf("OutputProfileSaveError: %v\n", err)
	return nil
}

func (pp *ProfilePresenter) OutputSuccessUpdate() error {
	fmt.Printf("Success Update Profile!")
	return nil
}

func (pp *ProfilePresenter) OutputProfile(profile *entities.Profile) error {
	fmt.Printf("Profile : %v\n", profile)
	return nil
}
