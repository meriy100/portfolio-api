package usecase

import (
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type SkillInteractor struct {
	outputPort      ports.SkillOutputPort
	postRepository  ports.PostRepository
	skillRepository ports.SkillRepository
}

func NewSkillInteractor(outputPort ports.SkillOutputPort, postRepository ports.PostRepository, skillRepository ports.SkillRepository) ports.SkillInputPort {
	return &SkillInteractor{
		outputPort,
		postRepository,
		skillRepository,
	}
}

const SkillsPostID = 289

func (s *SkillInteractor) UpdateSkills() error {
	post, err := s.postRepository.FetchPost(SkillsPostID)
	if err != nil {
		return s.outputPort.OutputFetchPostError(err)
	}

	skills, err := post.ToSkills()
	if err != nil {
		return s.outputPort.OutputToSkillsError(err)
	}

	for _, skill := range skills {
		err := s.skillRepository.Save(skill)
		if err != nil {
			return s.outputPort.OutputSkillSaveError(skill, err)
		}
	}

	return s.outputPort.OutputSuccessUpdate()
}

func (s *SkillInteractor) IndexSkills() error {
	skills, err := s.skillRepository.All()
	if err != nil {
		return s.outputPort.OutputFetchSkillsError(err)
	}
	return s.outputPort.OutputSkills(skills)
}
