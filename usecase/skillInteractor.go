package usecase

import (
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type SkillInteractor struct {
	outputPort      ports.SkillOutputPort
	postRepository  ports.PostRepository
	SkillRepository ports.SkillRepository
}

func NewSkillInteractor(outputPort ports.SkillOutputPort, postRepository ports.PostRepository, skillRepository ports.SkillRepository) ports.SkillInputPort {
	return &SkillInteractor{
		outputPort,
		postRepository,
		skillRepository,
	}
}

func (s *SkillInteractor) UpdateSkills() error {
	post, err := s.postRepository.FetchPost(289)
	if err != nil {
		return s.outputPort.OutputFetchPostError(err)
	}

	skills, err := entities.ToSkills(post)
	if err != nil {
		return s.outputPort.OutputToSkillsError(err)
	}

	for _, skill := range skills {
		err := s.SkillRepository.Save(skill)
		if err != nil {
			return s.outputPort.OutputSkillSaveError(skill, err)
		}
	}

	return s.outputPort.OutputSuccessUpdate()
}

func (s *SkillInteractor) IndexSkills() error {
	skills, err := s.SkillRepository.All()
	if err != nil {
		return s.outputPort.OutputFetchSkillsError(err)
	}
	return s.outputPort.OutputSkills(skills)
}
