package entities

type History struct {
	Organization string    `json:"organization"`
	Products     []Product `json:"products"`
}

type Month struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type Product struct {
	Title        string   `json:"title"`
	StartMonth   Month    `json:"startMonth"`
	EndMonth     *Month   `json:"endMonth"`
	Description  []string `json:"description"`
	Technologies []string `json:"technologies"`
}

func CompareMonth(x, y Month) int {
	if x.Year == y.Year {
		if x.Month < y.Month {
			return -1
		} else {
			return 1
		}
	}
	if x.Year < y.Year {
		return -1
	}
	return 1
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
