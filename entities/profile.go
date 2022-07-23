package entities

import "time"

type Profile struct {
	Job         string
	Description string
	timestamp   time.Time
}

func NewProfile(job string, description string) *Profile {
	return &Profile{job, description, time.Now()}
}
