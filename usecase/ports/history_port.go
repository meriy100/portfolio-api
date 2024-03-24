package ports

import "github.com/meriy100/portfolio-api/entities"

type HistoryRepository interface {
	Save(history *entities.History) error
	All() ([]*entities.History, error)
}

type HistoryInputPort interface {
	UpdateHistories() error
	IndexHistories() error
}

type HistoryOutputPort interface {
	OutputFetchPostError(error) error
	OutputToHistoriesError(error) error
	OutputHistorySaveError(*entities.History, error) error
	OutputFetchHistoriesError(error) error
	OutputDeployError(error) error

	OutputSuccessUpdate() error
	OutputHistories([]*entities.History) error
}
