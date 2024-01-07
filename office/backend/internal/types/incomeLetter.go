package types

type IncomeLetterViewModel struct {
	ID          string `json:"id"`
	Number      uint   `json:"number"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	SubjectId   string `json:"subject"`
	Created_At  string `json:"created_at"`
	Owner       string `json:"owner"`
	Destination string `json:"destination"`
	OperatorId  string `json:"operatorId"`
	Status      uint   `json:"status"`
}

type IncomeLetterDTO struct {
	Number      uint   `json:"number"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	SubjectId   string `json:"subject"`
	Owner       string `json:"owner"`
	Destination string `json:"destination"`
	OperatorId  string `json:"operatorId"`
	Status      uint   `json:"status"`
}
