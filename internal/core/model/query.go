package model

const (
	defaultPage    = 0
	defaultPerPage = 10
)

type Query struct {
	Page    int    `form:"page" json:"page"`
	PerPage int    `form:"per_page" json:"per_page"`
	Order   string `form:"order" json:"order"`
	Keyword string `form:"query" json:"query"`
}

func NewQuery() *Query {
	return &Query{
		Page:    defaultPage,
		PerPage: defaultPerPage,
		Order:   "",
		Keyword: "",
	}
}
