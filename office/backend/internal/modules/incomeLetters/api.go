package incomeletters

import (
	"errors"
	"net/http"

	"mehrangcode.ir/office/internal/storage"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/utils"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	repo := storage.NewIncomSqliteRepository()
	letters, err := repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	statusCode := http.StatusOK
	if letters == nil {
		statusCode = http.StatusNoContent
	}

	utils.WriteJson(w, statusCode, letters)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var u types.IncomeLetterDTO
	var err error
	err = utils.ReadJson(w, r, &u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if u.Number == 0 {
		utils.ResponseToError(w, errors.New("number fields is required"), http.StatusBadRequest)
		return
	}
	repo := storage.NewIncomSqliteRepository()
	letterId, err := repo.Create(u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"letterId": letterId,
	})
}
