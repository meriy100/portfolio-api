package entities

import "time"

type Profile struct {
	Job         string    `json:"job"`
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

func NewProfile(job string, description string) *Profile {
	return &Profile{job, description, time.Now()}
}
