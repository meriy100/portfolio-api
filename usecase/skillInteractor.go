package usecase

import (
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
	"time"
)

type SkillInteractor struct {
	outputPort        ports.SkillOutputPort
	postRepository    ports.PostRepository
	historyRepository ports.HistoryRepository
	skillRepository   ports.SkillRepository
}

func NewSkillInteractor(outputPort ports.SkillOutputPort, postRepository ports.PostRepository, historyRepository ports.HistoryRepository, skillRepository ports.SkillRepository) ports.SkillInputPort {
	return &SkillInteractor{
		outputPort,
		postRepository,
		historyRepository,
		skillRepository,
	}
}

func (s *SkillInteractor) UpdateSkills() error {
	post, err := s.postRepository.FetchPost(289)
	if err != nil {
		return s.outputPort.OutputFetchPostError(err)
	}

	skills, err := post.ToSkills()
	if err != nil {
		return s.outputPort.OutputToSkillsError(err)
	}

	histories, err := s.historyRepository.All()
	if err != nil {
		return s.outputPort.OutputFetchSkillsError(err)
	}

	skillMap := entities.SkillMap{}

	for _, history := range histories {
		skillMap = history.SkillMap(skillMap)
	}

	now := entities.Month{Year: time.Now().Year(), Month: int(time.Now().Month())}

	for _, skill := range skills {
		skill.DurationMonth = entities.SumTerm(now, skillMap[skill.Name])
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
