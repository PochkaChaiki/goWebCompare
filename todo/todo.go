package todo

type Todo struct {
	Id        int    `json:"id,omitempty"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
