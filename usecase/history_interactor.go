package usecase

import (
	"context"

	"github.com/meriy100/portfolio-api/usecase/ports"
)

type HistoryInteractor struct {
	outputPort                ports.HistoryOutputPort
	postRepository            ports.PostRepository
	HistoryRepository         ports.HistoryRepository
	contentDeliveryRepository ports.ContentDeliveryRepository
}

func NewHistoryInteractor(outputPort ports.HistoryOutputPort, postRepository ports.PostRepository, profileRepository ports.HistoryRepository, contentDeliveryRepository ports.ContentDeliveryRepository) ports.HistoryInputPort {
	return &HistoryInteractor{
		outputPort,
		postRepository,
		profileRepository,
		contentDeliveryRepository,
	}
}

const HistoriesPostID = 254

func (h *HistoryInteractor) UpdateHistories() error {
	post, err := h.postRepository.FetchPost(HistoriesPostID)
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

	if err := h.contentDeliveryRepository.Deploy(context.Background()); err != nil {
		return h.outputPort.OutputDeployError(err)
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
