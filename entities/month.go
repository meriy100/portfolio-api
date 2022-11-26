package entities

type Month struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

func subMonth(x, y Month) int {
	return (x.Year-y.Year)*12 + (x.Month - y.Month) + 1
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
