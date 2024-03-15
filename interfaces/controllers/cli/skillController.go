package cli

import (
	"fmt"

	"github.com/meriy100/portfolio-api/usecase/ports"
)

type skillInputPortFactory func(ports.SkillOutputPort, ports.PostRepository, ports.SkillRepository) ports.SkillInputPort
type skillOutputPortFactory func() ports.SkillOutputPort

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

func (s *SkillController) UpdateSkills() {
	if err := s.newInputPort().UpdateSkills(); err != nil {
		fmt.Println(err)
	}
}

func (s *SkillController) IndexSkills() {
	if err := s.newInputPort().IndexSkills(); err != nil {
		fmt.Println(err)
	}
}

func (s *SkillController) newInputPort() ports.SkillInputPort {
	return s.InputPortFactory(s.OutputPortFactory(), s.PostRepository, s.SkillRepository)
}
