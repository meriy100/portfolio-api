package cli

import (
	"fmt"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type historyInputPortFactory func(ports.HistoryOutputPort, ports.PostRepository, ports.HistoryRepository) ports.HistoryInputPort
type historyOutputPortFactory func() ports.HistoryOutputPort

type HistoryController struct {
	HistoryRepository ports.HistoryRepository
	PostRepository    ports.PostRepository
	InputPortFactory  historyInputPortFactory
	OutputPortFactory historyOutputPortFactory
}

func NewHistoryController(
	historyRepository ports.HistoryRepository,
	postRepository ports.PostRepository,
	inputFactory historyInputPortFactory,
	outputFactory historyOutputPortFactory,
) *HistoryController {
	return &HistoryController{
		historyRepository,
		postRepository,
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
	return h.InputPortFactory(h.OutputPortFactory(), h.PostRepository, h.HistoryRepository)
}
