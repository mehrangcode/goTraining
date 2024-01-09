package types

type SubjectViewModel struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Archive bool   `json:"archive"`
}

type SubjectDTO struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Archive bool   `json:"archive"`
}
