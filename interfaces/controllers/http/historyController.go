package http

import (
	"fmt"
	"github.com/meriy100/portfolio-api/usecase/ports"
	"net/http"
)

type historyInputPortFactory func(ports.HistoryOutputPort, ports.PostRepository, ports.HistoryRepository) ports.HistoryInputPort
type historyOutputPortFactory func(http.ResponseWriter) ports.HistoryOutputPort

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
func (h *HistoryController) UpdateHistories(w http.ResponseWriter, r *http.Request) {
	if err := h.newInputPort(w).UpdateHistories(); err != nil {
		fmt.Println(err)
	}
}

func (h *HistoryController) IndexHistories(w http.ResponseWriter, r *http.Request) {
	if err := h.newInputPort(w).IndexHistories(); err != nil {
		fmt.Println(err)
	}
}

func (h *HistoryController) newInputPort(w http.ResponseWriter) ports.HistoryInputPort {
	return h.InputPortFactory(h.OutputPortFactory(w), h.PostRepository, h.HistoryRepository)
}
