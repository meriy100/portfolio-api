package http

import (
	"fmt"
	"net/http"

	"github.com/meriy100/portfolio-api/usecase/ports"
)

type historyInputPortFactory func(ports.HistoryOutputPort, ports.PostRepository, ports.HistoryRepository, ports.ContentDeliveryRepository) ports.HistoryInputPort
type historyOutputPortFactory func(http.ResponseWriter) ports.HistoryOutputPort

type HistoryController struct {
	HistoryRepository         ports.HistoryRepository
	PostRepository            ports.PostRepository
	ContentDeliveryRepository ports.ContentDeliveryRepository
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
	return h.InputPortFactory(h.OutputPortFactory(w), h.PostRepository, h.HistoryRepository, h.ContentDeliveryRepository)
}
