package types

type IssuedLetterViewModel struct {
	ID          string `json:"id"`
	Number      uint   `json:"number"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	SubjectId   string `json:"subjectId" db:"subjectId"`
	Created_At  string `json:"created_at" db:"created_at"`
	Owner       string `json:"owner"`
	Destination string `json:"destination"`
	OperatorId  string `json:"operatorId" db:"operatorId"`
	Status      uint   `json:"status"`
}

type IssuedLetterDTO struct {
	Number      uint   `json:"number"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	SubjectId   string `json:"subjectId"`
	Owner       string `json:"owner"`
	Destination string `json:"destination"`
	OperatorId  string `json:"operatorId"`
	Status      uint   `json:"status"`
}
