package types

type IssuedLetterViewModel struct {
	ID           string  `json:"id"`
	Number       uint    `json:"number"`
	Title        string  `json:"title"`
	Content      string  `json:"content"`
	SubjectId    string  `json:"subjectId" db:"subjectId"`
	SubjectName  *string `json:"subjectName" db:"subjectName"`
	Created_at   string  `json:"created_at" db:"created_at"`
	Owner        string  `json:"owner"`
	Destination  string  `json:"destination"`
	OperatorId   string  `json:"operatorId" db:"operatorId"`
	OperatorName *string `json:"operatorName" db:"operatorName"`
	Status       uint    `json:"status"`
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
