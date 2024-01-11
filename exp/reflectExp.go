package main

import (
	"database/sql"
	"fmt"
)

type IssuedLetterViewModel struct {
	Title       string         `json:"title"`
	SubjectName sql.NullString `json:"subjectName" db:"subjectName"`
}

func checkReflet() {

	fmt.Println("DDD")
}
