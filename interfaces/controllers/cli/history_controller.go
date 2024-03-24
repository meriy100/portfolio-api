package cli

import (
	"fmt"

	"github.com/meriy100/portfolio-api/usecase/ports"
)

type historyInputPortFactory func(ports.HistoryOutputPort, ports.PostRepository, ports.HistoryRepository, ports.ContentDeliveryRepository) ports.HistoryInputPort
type historyOutputPortFactory func() ports.HistoryOutputPort

type HistoryController struct {
	HistoryRepository         ports.HistoryRepository
	PostRepository            ports.PostRepository
	contentDeliveryRepository ports.ContentDeliveryRepository
	InputPortFactory          historyInputPortFactory
	OutputPortFactory         historyOutputPortFactory
}

func NewHistoryController(
	historyRepository ports.HistoryRepository,
	postRepository ports.PostRepository,
	contentDeliveryRepository ports.ContentDeliveryRepository,
	inputFactory historyInputPortFactory,
	outputFactory historyOutputPortFactory,
) *HistoryController {
	return &HistoryController{
		historyRepository,
		postRepository,
		contentDeliveryRepository,
		inputFactory,
		outputFactory,
	}
}

func (h *HistoryController) UpdateHistories() {
	if err := h.newInputPort().UpdateHistories(); err != nil {
		fmt.Println(err)
	}
}

func (h *HistoryController) IndexHistories() {
	if err := h.newInputPort().IndexHistories(); err != nil {
		fmt.Println(err)
	}
}

func (h *HistoryController) newInputPort() ports.HistoryInputPort {
	return h.InputPortFactory(h.OutputPortFactory(), h.PostRepository, h.HistoryRepository, h.contentDeliveryRepository)
}
