package cli

import (
	"fmt"

	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type SkillPresenter struct {
}

func NewSkillPresenter() ports.SkillOutputPort {
	return &SkillPresenter{}
}

func (sp *SkillPresenter) OutputFetchPostError(err error) error {
	fmt.Printf("OutputFetchPostError: %v\n", err)
	return nil
}

func (sp *SkillPresenter) OutputToSkillsError(err error) error {
	fmt.Printf("OutputToSkillError: %v\n", err)
	return nil
}

func (sp *SkillPresenter) OutputSkillSaveError(skill *entities.Skill, err error) error {
	fmt.Printf("OutputSkillSaveError: %v, %v\n", skill, err)
	return nil
}

func (sp *SkillPresenter) OutputFetchSkillsError(err error) error {
	fmt.Printf("OutputFetchSkillsError: %v\n", err)
	return nil
}

func (sp *SkillPresenter) OutputSuccessUpdate() error {
	fmt.Printf("Success Update Skill!\n")
	return nil
}

func (sp *SkillPresenter) OutputSkills(skills []*entities.Skill) error {
	fmt.Printf("Skill:\n")
	for _, skill := range skills {
		fmt.Printf("%v\n", skill)
	}
	return nil
}
