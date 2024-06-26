package cli

import (
	"fmt"

	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type HistoryPresenter struct {
}

func NewHistoryPresenter() ports.HistoryOutputPort {
	return &HistoryPresenter{}
}

func (h *HistoryPresenter) OutputFetchPostError(err error) error {
	fmt.Printf("OutputFetchPostError: %v\n", err)
	return nil
}

func (h *HistoryPresenter) OutputToHistoriesError(err error) error {
	fmt.Printf("OutputToHistoriesError: %v\n", err)
	return nil
}

func (h *HistoryPresenter) OutputHistorySaveError(history *entities.History, err error) error {
	fmt.Printf("OutputHistorySaveError: %v,  %v\n", history.Organization, err)
	return nil
}

func (h *HistoryPresenter) OutputFetchHistoriesError(err error) error {
	fmt.Printf("OutputFetchHistoriesError: %v\n", err)
	return nil
}

func (h *HistoryPresenter) OutputDeployError(err error) error {
	fmt.Printf("OutputDeployError: %v\n", err)
	return nil
}

func (h *HistoryPresenter) OutputSuccessUpdate() error {
	fmt.Printf("Success Update Histories!")
	return nil
}
func (h *HistoryPresenter) OutputHistories(histories []*entities.History) error {
	fmt.Printf("Histories:\n")
	for _, history := range histories {
		fmt.Printf("%v\n", history)
	}
	return nil
}
