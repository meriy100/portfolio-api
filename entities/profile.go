package entities

import "time"

type Profile struct {
	Job              string    `json:"job"`
	Description      string    `json:"description"`
	SkillDescription []string  `json:"skillDescription"`
	Licenses         []string  `json:"licenses"`
	Pr               string    `json:"pr"`
	Timestamp        time.Time `json:"timestamp"`
}

func NewProfile(job, description string, skillDescription, licenses []string, pr string) *Profile {
	return &Profile{job, description, skillDescription, licenses, pr, time.Now()}
}
