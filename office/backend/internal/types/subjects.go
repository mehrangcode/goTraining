package types

type SubjectViewModel struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Archive uint16 `json:"archive"`
}

type SubjectDTO struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Archive uint16 `json:"archive"`
}
