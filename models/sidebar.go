package models

type Sidebar struct {
	Id          int
	TabName     string
	FontClass   string
	UrlPath     string
}

func (m *Sidebar) TableName() string {
	return TableName("sidebar")
}
