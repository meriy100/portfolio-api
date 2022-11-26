package entities

import (
	"fmt"
	"strings"
	"time"
)

type SkillCategory int

const (
	Os SkillCategory = iota
	Language
	FrameworkLibrary
	Db
	CiCd
	Infrastructure
	Other
)

type Skill struct {
	Name        string
	Lv          int
	Description string
	Category    SkillCategory
	Timestamp   time.Time `json:"timestamp"`
}

func NewSkill(name string, lv int, description string, category SkillCategory) *Skill {
	return &Skill{
		name,
		lv,
		description,
		category,
		time.Now(),
	}
}

func skillCategoryDecoder(text string) (SkillCategory, error) {
	switch strings.TrimSpace(text) {
	case "os":
		return Os, nil
	case "language":
		return Language, nil
	case "frameworkLibrary":
		return FrameworkLibrary, nil
	case "db":
		return Db, nil
	case "cicd":
		return CiCd, nil
	case "infrastructure":
		return Infrastructure, nil
	case "other":
		return Other, nil
	default:
		return SkillCategory(-1), fmt.Errorf("can't decode skillCateogry part. part is '%s'", strings.TrimSpace(text))
	}
}
