package ports

import "github.com/meriy100/portfolio-api/entities"

type SkillRepository interface {
	Save(skill *entities.Skill) error
	All() ([]*entities.Skill, error)
}

type SkillInputPort interface {
	UpdateSkills() error
	IndexSkills() error
}

type SkillOutputPort interface {
	OutputFetchPostError(error) error
	OutputToSkillsError(error) error
	OutputSkillSaveError(*entities.Skill, error) error
	OutputFetchSkillsError(error) error

	OutputSuccessUpdate() error
	OutputSkills([]*entities.Skill) error
}
