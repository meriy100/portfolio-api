package entities

import "time"

type Profile struct {
	Job              string    `json:"job"`
	Description      string    `json:"description"`
	SkillDescription []string  `json:"skillDescription"`
	Timestamp        time.Time `json:"timestamp"`
}

func NewProfile(job string, description string, skillDescription []string) *Profile {
	return &Profile{job, description, skillDescription, time.Now()}
}
