package dbcon

type Student struct {
	ID       int    `json:"ID,omitempty"`
	Name     string `json:"name,omitempty"`
	Course   string `json:"course,omitempty"`
	Fees     int    `json:"fees,omitempty"`
	Contact  string `json:"contact,omitempty"`
	Subjects string `json:"subjects,omitempty"`
}
