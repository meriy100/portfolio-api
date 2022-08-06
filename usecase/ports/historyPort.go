package ports

import "github.com/meriy100/portfolio-api/entities"

type HistoryRepository interface {
	Save(history *entities.History) error
}

type HistoryInputPort interface {
	UpdateHistories() error
}

type HistoryOutputPort interface {
	OutputFetchPostError(error) error
	OutputToHistoriesError(error) error
	OutputHistorySaveError(*entities.History, error) error

	OutputSuccessUpdate() error
}
