package presenters

import (
	"fmt"
	"github.com/meriy100/portfolio-api/usecase/ports"
	"net/http"
)

type ProfileHttpPresenter struct {
	w http.ResponseWriter
}

func NewProfileHttpPresenter(w http.ResponseWriter) ports.ProfileOutputPort {
	return &ProfileHttpPresenter{w}
}

func (pp *ProfileHttpPresenter) OutputFetchPostError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputFetchPostError: %v", err.Error()), http.StatusInternalServerError)

	return nil
}

func (pp *ProfileHttpPresenter) OutputToProfileError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputToProfileError: %v\n", err.Error()), http.StatusInternalServerError)
	return nil
}

func (pp *ProfileHttpPresenter) OutputProfileSaveError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputProfileSaveError: %v\n", err.Error()), http.StatusInternalServerError)
	return nil
}

func (pp *ProfileHttpPresenter) OutputSuccessUpdate() error {
	_, err := fmt.Fprintf(pp.w, "Success Update Profile!")
	return err
}
