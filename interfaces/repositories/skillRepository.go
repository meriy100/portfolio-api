package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/meriy100/portfolio-api/entities"
	"google.golang.org/api/iterator"
)

type SkillRepository struct {
	ctx    context.Context
	client *firestore.Client
}

func NewSkillRepository(ctx context.Context, client *firestore.Client) *SkillRepository {
	return &SkillRepository{ctx, client}
}

func (s *SkillRepository) Save(skill *entities.Skill) error {
	_, err := s.client.Collection("portfolio-data-skills").Doc(skill.Name).Set(s.ctx, skill)
	if err != nil {
		return err
	}
	return nil
}

func (s *SkillRepository) All() ([]*entities.Skill, error) {
	var skills []*entities.Skill
	iter := s.client.Collection("portfolio-data-skills").Documents(s.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return skills, err
		}

		var skill entities.Skill
		err = doc.DataTo(&skill)
		if err != nil {
			return skills, err
		}
		skills = append(skills, &skill)
	}

	return skills, nil
}
