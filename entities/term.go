package entities

type Term struct {
	StartMonth Month
	EndMonth   Month
}

func ToTerm(startMonth Month, endMonth *Month, now Month) Term {
	if endMonth != nil {
		return Term{startMonth, *endMonth}
	}
	return Term{startMonth, now}
}

func SumTerm(terms []Term) int {
	months := map[Month]bool{}
	for _, term := range terms {
		for _, month := range rangeMonths(terms) {
			months[month] = true
		}
	}
	return len(months)
}

func rangeMonths(term Term) []Month {
	ms := []Month{}
	for m := term.StartMonth; m == term.EndMonth; m = nextMonth(term.StartMonth) {
		ms = append(ms, m)
	}
	return ms
}
