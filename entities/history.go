package entities

type History struct {
	Organization string    `json:"organization"`
	Products     []Product `json:"products"`
}

type Product struct {
	Title        string   `json:"title"`
	StartMonth   Month    `json:"startMonth"`
	EndMonth     *Month   `json:"endMonth"`
	Description  []string `json:"description"`
	Technologies []string `json:"technologies"`
}

func (h *History) StartMonth() Month {
	var startMonth Month
	for _, product := range h.Products {
		if startMonth.Year == 0 {
			startMonth = product.StartMonth
		}
		if CompareMonth(product.StartMonth, startMonth) == -1 {
			startMonth = product.StartMonth
		}
	}
	return startMonth
}

func (h *History) EndMonth() *Month {
	var endMonth *Month
	for _, product := range h.Products {
		if product.EndMonth != nil {
			if endMonth == nil {
				endMonth = product.EndMonth
			}
			if CompareMonth(*product.EndMonth, *endMonth) == 1 {
				endMonth = product.EndMonth
			}
		}
	}
	return endMonth
}

type SkillMap map[string][]Term

func (h *History) SkillMap(skillMap SkillMap) SkillMap {
	for _, product := range h.Products {
		for _, technology := range product.Technologies {
			if skillMap[technology] == nil {
				skillMap[technology] = []Term{}
			}
			skillMap[technology] = append(skillMap[technology], Term{product.StartMonth, product.EndMonth})
		}
	}
	return skillMap
}
