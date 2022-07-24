package entities

type History struct {
	Organization string
}

type Month struct {
	Year  int
	Month int
}

type Product struct {
	Title        string
	StartMonth   Month
	EndMonth     Month
	Description  string
	Technologies []string
}
