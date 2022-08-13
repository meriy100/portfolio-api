package http

import (
	"encoding/json"
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
	"net/http"
)

type ProfilePresenter struct {
	w http.ResponseWriter
}

func NewProfilePresenter(w http.ResponseWriter) ports.ProfileOutputPort {
	return &ProfilePresenter{w}
}

func (pp *ProfilePresenter) OutputFetchPostError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputFetchPostError: %v", err.Error()), http.StatusInternalServerError)

	return nil
}

func (pp *ProfilePresenter) OutputFindProfileError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputFindProfileError: %v", err.Error()), http.StatusInternalServerError)

	return nil
}

func (pp *ProfilePresenter) OutputToProfileError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputToProfileError: %v\n", err.Error()), http.StatusInternalServerError)
	return nil
}

func (pp *ProfilePresenter) OutputProfileSaveError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputProfileSaveError: %v\n", err.Error()), http.StatusInternalServerError)
	return nil
}

func (pp *ProfilePresenter) OutputSuccessUpdate() error {
	_, err := fmt.Fprintf(pp.w, "Success Update Profile!")
	return err
}

func (pp *ProfilePresenter) OutputProfile(profile *entities.Profile) error {
	j, err := json.Marshal(ResponseData{profile})
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(pp.w, string(j))
	return err
}
