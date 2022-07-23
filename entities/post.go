package entities

import "strings"

type Post struct {
	BodyMd string `json:"body_md"`
}

func (p *Post) ToProfile() (*Profile, error) {
	parts := strings.Split(p.BodyMd, "##")

	var job string
	var description string

	for _, part := range parts {
		if strings.HasPrefix(part, " job") {
			job = strings.Replace(part, "job\r\n", "", 1)
		}

		if strings.HasPrefix(part, " description") {
			description = strings.Replace(part, "description\r\n", "", 1)
		}

	}

	return NewProfile(job, description), nil
}
