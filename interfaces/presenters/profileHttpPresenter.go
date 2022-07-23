package presenters

import (
	"encoding/json"
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
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

func (pp *ProfileHttpPresenter) OutputFindProfileError(err error) error {
	http.Error(pp.w, fmt.Sprintf("OutputFindProfileError: %v", err.Error()), http.StatusInternalServerError)

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

func (pp *ProfileHttpPresenter) OutputProfile(profile *entities.Profile) error {
	j, err := json.Marshal(profile)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(pp.w, string(j))
	return err
}
