package http

import (
	"fmt"
	"net/http"

	"github.com/meriy100/portfolio-api/usecase/ports"
)

type skillInputPortFactory func(ports.SkillOutputPort, ports.PostRepository, ports.SkillRepository) ports.SkillInputPort
type skillOutputPortFactory func(http.ResponseWriter) ports.SkillOutputPort

type SkillController struct {
	SkillRepository   ports.SkillRepository
	PostRepository    ports.PostRepository
	InputPortFactory  skillInputPortFactory
	OutputPortFactory skillOutputPortFactory
}

func NewSkillController(
	skillRepository ports.SkillRepository,
	postRepository ports.PostRepository,
	inputFactory skillInputPortFactory,
	outputFactory skillOutputPortFactory,
) *SkillController {
	return &SkillController{
		skillRepository,
		postRepository,
		inputFactory,
		outputFactory,
	}
}
func (h *SkillController) UpdateSkills(w http.ResponseWriter, r *http.Request) {
	if err := h.newInputPort(w).UpdateSkills(); err != nil {
		fmt.Println(err)
	}
}

func (h *SkillController) IndexSkills(w http.ResponseWriter, r *http.Request) {
	if err := h.newInputPort(w).IndexSkills(); err != nil {
		fmt.Println(err)
	}
}

func (h *SkillController) newInputPort(w http.ResponseWriter) ports.SkillInputPort {
	return h.InputPortFactory(h.OutputPortFactory(w), h.PostRepository, h.SkillRepository)
}
