package usecase

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/Songmu/flextime"
	"github.com/google/go-cmp/cmp"
	"github.com/meriy100/portfolio-api/entities"
	"github.com/meriy100/portfolio-api/usecase/ports"
)

type TestSkillOutputPort struct {
	ports.SkillOutputPort
}

func (tSOP *TestSkillOutputPort) OutputSuccessUpdate() error {
	return nil
}

type TestPostRepository struct {
	ports.PostRepository
	post *entities.Post
}

func (tPR *TestPostRepository) FetchPost(postId int) (*entities.Post, error) {
	if postId != 289 {
		return nil, fmt.Errorf("postId want 289 but %v", postId)
	}
	return tPR.post, nil
}

type TestSkillRepository struct {
	ports.SkillRepository
	insert []*entities.Skill
	all    []*entities.Skill
}

func (tSR *TestSkillRepository) Save(skill *entities.Skill) error {
	if tSR.insert == nil {
		tSR.insert = []*entities.Skill{}
	}
	tSR.insert = append(tSR.insert, skill)
	return nil
}

func (tSR *TestSkillRepository) All() ([]*entities.Skill, error) {
	return tSR.all, nil
}

func TestNewSkillInteractor(t *testing.T) {
	type args struct {
		outputPort      ports.SkillOutputPort
		postRepository  ports.PostRepository
		skillRepository ports.SkillRepository
	}
	tests := []struct {
		name string
		args args
		want ports.SkillInputPort
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSkillInteractor(tt.args.outputPort, tt.args.postRepository, tt.args.skillRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSkillInteractor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSkillInteractor_IndexSkills(t *testing.T) {
	type fields struct {
		outputPort      ports.SkillOutputPort
		postRepository  ports.PostRepository
		skillRepository ports.SkillRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SkillInteractor{
				outputPort:      tt.fields.outputPort,
				postRepository:  tt.fields.postRepository,
				skillRepository: tt.fields.skillRepository,
			}
			if err := s.IndexSkills(); (err != nil) != tt.wantErr {
				t.Errorf("IndexSkills() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSkillInteractor_UpdateSkills(t *testing.T) {
	now := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	type fields struct {
		outputPort      ports.SkillOutputPort
		postRepository  ports.PostRepository
		skillRepository *TestSkillRepository
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*entities.Skill
		wantErr bool
	}{
		{
			"success",
			fields{
				&TestSkillOutputPort{},
				&TestPostRepository{
					post: &entities.Post{BodyMd: "# os\n## test\n### description\ntestD\n### lv\n2\n"},
				},
				&TestSkillRepository{},
			},
			[]*entities.Skill{
				{
					Name:        "test",
					Lv:          2,
					Description: "testD",
					Category:    entities.Os,
					Timestamp:   now,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flextime.Fix(now)
			defer flextime.Restore()

			s := &SkillInteractor{
				outputPort:      tt.fields.outputPort,
				postRepository:  tt.fields.postRepository,
				skillRepository: tt.fields.skillRepository,
			}
			if err := s.UpdateSkills(); (err != nil) != tt.wantErr {
				t.Errorf("UpdateSkills() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(tt.want) != len(tt.fields.skillRepository.insert) {
				t.Errorf("UdateSkills() saved value = %v, want %v", tt.fields.skillRepository.insert, tt.want)
			}
			if diff := cmp.Diff(tt.want, tt.fields.skillRepository.insert); diff != "" {
				t.Errorf("UpdateSkills() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
