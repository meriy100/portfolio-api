package entities

type History struct {
	Organization string
	Products     []Product
}

type Month struct {
	Year  int
	Month int
}

type Product struct {
	Title        string
	StartMonth   Month
	EndMonth     *Month
	Description  []string
	Technologies []string
}
