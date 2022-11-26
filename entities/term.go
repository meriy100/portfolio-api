package entities

type Term struct {
	StartMonth Month
	EndMonth   *Month
}

func SumTerm(now Month, terms []Term) int {
	month := 0
	for _, term := range terms {
		if term.EndMonth != nil {
			month = month + subMonth(*term.EndMonth, term.StartMonth)
		} else {
			month = month + subMonth(now, term.StartMonth)
		}
	}
	return month
}
