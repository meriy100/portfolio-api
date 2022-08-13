package http

import (
	"encoding/json"
	"fmt"
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
	"net/http"
)

type HistoryPresenter struct {
	w http.ResponseWriter
}

func NewHistoryPresenter(w http.ResponseWriter) ports.HistoryOutputPort {
	return &HistoryPresenter{w}
}

func (h *HistoryPresenter) OutputFetchPostError(err error) error {
	http.Error(h.w, fmt.Sprintf("OutputFetchPostError: %v\n", err), http.StatusInternalServerError)
	return nil
}

func (h *HistoryPresenter) OutputToHistoriesError(err error) error {
	http.Error(h.w, fmt.Sprintf("OutputToHistoriesError: %v\n", err), http.StatusInternalServerError)
	return nil
}

func (h *HistoryPresenter) OutputHistorySaveError(history *entities.History, err error) error {
	http.Error(h.w, fmt.Sprintf("OutputHistorySaveError: %v,  %v\n", history.Organization, err), http.StatusInternalServerError)
	return nil
}

func (h *HistoryPresenter) OutputFetchHistoriesError(err error) error {
	http.Error(h.w, fmt.Sprintf("OutputFetchHistoriesError: %v\n", err), http.StatusInternalServerError)
	return nil
}

func (h *HistoryPresenter) OutputSuccessUpdate() error {
	_, err := fmt.Fprintf(h.w, "Success Update Histories!")
	return err
}
func (h *HistoryPresenter) OutputHistories(histories []*entities.History) error {
	j, err := json.Marshal(ResponseData{histories})
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(h.w, string(j))

	return err
}
