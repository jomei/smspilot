package templateDB

type TemplateDB struct {
	Id     int    `json:"id,omitempty"`
	Text   string `json:"text,omitempty"`
	Status int    `json:"status,omitempty"`
}
