package types

type LetterTemplateViewModel struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Content     string  `json:"content"`
	Status      uint    `json:"status"`
	SubjectId   *string `json:"subjectId" db:"subjectId"`
	SubjectName *string `json:"subjectName" db:"subjectName"`
}

type LetterTemplateDTO struct {
	Title     string  `json:"title"`
	Content   string  `json:"content"`
	Status    uint    `json:"status"`
	SubjectId *string `json:"subjectId" db:"subjectId"`
}
