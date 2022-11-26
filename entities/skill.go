package entities

import (
	"fmt"
	"strconv"
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

func ToSkills(post *Post) ([]*Skill, error) {
	var skills []*Skill
	parts := splitLv1(post.BodyMd)
	for _, part := range parts {
		head, body := separateHeadTail(part)
		category, err := skillCategoryDecoder(head)
		if err != nil {
			return skills, err
		}
		for _, skillPart := range splitLv2(body) {
			skill, err := toSkill(skillPart, category)
			if err != nil {
				return skills, err
			}

			skills = append(skills, skill)
		}

	}
	return skills, nil
}

func toSkill(part string, category SkillCategory) (*Skill, error) {
	var lv int
	var description string
	head, body := separateHeadTail(part)

	name := strings.TrimSpace(head)

	for _, p := range splitLv3(body) {
		key, text := separateHeadTail(p)
		switch key {
		case "lv":
			i, err := strconv.Atoi(strings.TrimSpace(text))
			if err != nil {
				return &Skill{}, fmt.Errorf("%s's lv decode error. %s", name, err.Error())
			}

			lv = i
		case "description":
			description = compactNl(text)
		default:
			fmt.Printf("key: %v\n", key)
		}
	}

	return NewSkill(name, lv, description, category), nil
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
