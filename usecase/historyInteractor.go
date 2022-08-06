package usecase

import "github.com/meriy100/portfolio-api/usecase/ports"

type HistoryInteractor struct {
	outputPort        ports.HistoryOutputPort
	postRepository    ports.PostRepository
	HistoryRepository ports.HistoryRepository
}

func NewHistoryInteractor(outputPort ports.HistoryOutputPort, postRepository ports.PostRepository, profileRepository ports.HistoryRepository) ports.HistoryInputPort {
	return &HistoryInteractor{
		outputPort,
		postRepository,
		profileRepository,
	}
}

func (h *HistoryInteractor) UpdateHistories() error {
	post, err := h.postRepository.FetchPost(254)
	if err != nil {
		return h.outputPort.OutputFetchPostError(err)
	}

	histories, err := post.ToHistories()
	if err != nil {
		return h.outputPort.OutputToHistoriesError(err)
	}

	for _, history := range histories {
		err := h.HistoryRepository.Save(history)
		if err != nil {
			return h.outputPort.OutputHistorySaveError(history, err)
		}
	}

	return h.outputPort.OutputSuccessUpdate()
}

func (h *HistoryInteractor) IndexHistories() error {
	histories, err := h.HistoryRepository.All()
	if err != nil {
		return h.outputPort.OutputFetchHistoriesError(err)
	}
	return h.outputPort.OutputHistories(histories)
}
