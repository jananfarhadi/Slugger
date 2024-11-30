package slugger

import "strings"

type NewSluggerInstance struct {
	Text *string `json:"text"`
	Slug *string `json:"slug"`
}

func NewSlug() *NewSluggerInstance {
	return &NewSluggerInstance{}
}

func (s *NewSluggerInstance) GetText(t string) {
	s.Text = &t
}

func (s *NewSluggerInstance) ToSlug() string {
	if s.Text == nil {
		return ""
	}

	str := strings.ToLower(*s.Text)
	str = strings.Replace(str, " ", "-", -1)
	s.Slug = &str

	return *s.Slug
}
