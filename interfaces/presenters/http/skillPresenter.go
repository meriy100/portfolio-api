package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type SkillPresenter struct {
	w http.ResponseWriter
}

func NewSkillPresenter(w http.ResponseWriter) ports.SkillOutputPort {
	return &SkillPresenter{w}
}

func (s *SkillPresenter) OutputFetchPostError(err error) error {
	http.Error(s.w, fmt.Sprintf("OutputFetchPostError: %v\n", err), http.StatusInternalServerError)
	return nil
}

func (s *SkillPresenter) OutputToSkillsError(err error) error {
	http.Error(s.w, fmt.Sprintf("OutputToSkillsError: %v\n", err), http.StatusInternalServerError)
	return nil
}

func (s *SkillPresenter) OutputSkillSaveError(skill *entities.Skill, err error) error {
	http.Error(s.w, fmt.Sprintf("OutputSkillSaveError: %v,  %v\n", skill.Name, err), http.StatusInternalServerError)
	return nil
}

func (s *SkillPresenter) OutputFetchSkillsError(err error) error {
	http.Error(s.w, fmt.Sprintf("OutputFetchSkillsError: %v\n", err), http.StatusInternalServerError)
	return nil
}

func (s *SkillPresenter) OutputSuccessUpdate() error {
	_, err := fmt.Fprintf(s.w, "Success Update Skills!")
	return err
}
func (s *SkillPresenter) OutputSkills(skills []*entities.Skill) error {
	sort.SliceStable(skills, func(x, y int) bool {
		return int(skills[x].Category) < int(skills[y].Category)
	})

	j, err := json.Marshal(ResponseData{skills})
	if err != nil {
		return err
	}

	_, err = fmt.Fprint(s.w, string(j))

	return err
}
