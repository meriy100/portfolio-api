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
