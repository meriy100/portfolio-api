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
	historyParts := splitLv1(p.BodyMd)

	for _, hp := range historyParts {
		history, err := toHistory(hp)
		if err != nil {
			return histories, err
		}
		histories = append(histories, history)
	}
	return histories, nil
}

func toHistory(historyPart string) (History, error) {
	history := History{}
	orgName, body := separateHeadTail(historyPart)
	history.Organization = orgName
	productParts := splitLv2(body)
	for _, pp := range productParts {
		product, err := toProduct(pp)
		if err != nil {
			return History{}, err
		}
		history.Products = append(history.Products, product)
	}
	return history, nil
}

func toProduct(productPart string) (Product, error) {
	product := Product{}
	prdName, prdBody := separateHeadTail(productPart)
	product.Title = prdName

	productColumns := splitLv3(prdBody)
	for _, pc := range productColumns {
		key, body := separateHeadTail(pc)
		switch key {
		case "startMonth":
			sm, err := strToMonth(body)
			if err != nil {
				return Product{}, err
			}
			product.StartMonth = sm
		case "endMonth":
			if len(body) != 0 {
				em, err := strToMonth(body)
				if err != nil {
					return Product{}, err
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
	return product, nil
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

func sliceCompact(xs []string) []string {
	var ss []string
	for _, s := range xs {
		if len(strings.TrimSpace(s)) > 0 {
			ss = append(ss, strings.TrimSpace(s))
		}
	}
	return ss
}

func splitLv1(s string) []string {
	return sliceCompact(regexp.MustCompile("(^|\n)#\\s+").Split(s, -1))
}

func splitLv2(s string) []string {
	return sliceCompact(regexp.MustCompile("(^|\n)##\\s+").Split(s, -1))
}

func splitLv3(s string) []string {
	return sliceCompact(regexp.MustCompile("(^|\n)###\\s+").Split(s, -1))
}
