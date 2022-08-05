package entities

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

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

func (p *Post) ToHistories() ([]History, error) {
	var histories []History
	organizationParts := regexp.MustCompile("(^|\n)#\\s+").Split(p.BodyMd, -1)

	for _, op := range organizationParts {
		if len(strings.TrimSpace(op)) > 0 {
			history := History{}
			orgName, orgBody := separateHeadTail(op)
			history.Organization = orgName
			productParts := regexp.MustCompile("(^|\n)##\\s+").Split(orgBody, -1)
			for _, pp := range productParts {
				if len(strings.TrimSpace(pp)) > 0 {
					product := Product{}
					prdName, prdBody := separateHeadTail(pp)
					product.Title = prdName

					productColumns := regexp.MustCompile("(^|\n)###\\s+").Split(prdBody, -1)
					for _, pc := range productColumns {
						if len(strings.TrimSpace(pc)) > 0 {
							key, body := separateHeadTail(pc)
							switch key {
							case "startMonth":
								sm, err := strToMonth(body)
								if err != nil {
									return []History{}, err
								}
								product.StartMonth = sm
							case "endMonth":
								if len(body) != 0 {
									em, err := strToMonth(body)
									if err != nil {
										return []History{}, err
									}
									product.EndMonth = &em
								}
							case "description":
								product.Description = mdListToSlice(body)
							case "technologyUsed":
								product.Technologies = mdListToSlice(body)
							default:
								fmt.Printf("key: %v\n", key)
							}
						}
					}
					history.Products = append(history.Products, product)
				}
			}
			histories = append(histories, history)
		}
	}
	return histories, nil
}

func separateHeadTail(s string) (string, string) {
	ss := strings.SplitN(s, "\r\n", 2)
	switch len(ss) {
	case 1:
		return strings.TrimSpace(ss[0]), ""
	case 2:
		return strings.TrimSpace(ss[0]), strings.TrimSpace(ss[1])
	default:
		return "", ""
	}
}

func strToMonth(s string) (Month, error) {
	ss := strings.Split(s, "/")
	if len(ss) != 2 {
		return Month{}, fmt.Errorf("can't month parse. string split '/', got %v", ss)
	}
	y, err := strconv.Atoi(ss[0])

	if err != nil {
		return Month{}, fmt.Errorf("can't year part to int, year part is %v", ss[0])
	}

	m, err := strconv.Atoi(ss[1])
	if err != nil {
		return Month{}, fmt.Errorf("can't month part to int, month part is %v", ss[1])
	}

	return Month{y, m}, nil
}

func mdListToSlice(s string) []string {
	var ss []string
	ssd := regexp.MustCompile("(^|\n)-\\s+").Split(s, -1)

	for _, sd := range ssd {
		tsd := strings.TrimSpace(sd)
		if len(tsd) > 0 {
			ss = append(ss, tsd)
		}
	}

	return ss
}
