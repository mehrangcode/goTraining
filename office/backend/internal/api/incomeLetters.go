package api

import (
	"errors"
	"net/http"

	"mehrangcode.ir/office/internal/storage"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/utils"
)

type IncomeLettersModule struct {
	repo *storage.IncomeLettersSqliteRepository
}

func NewIncomeLettersModule(repo storage.IncomeLettersSqliteRepository) *IncomeLettersModule {
	return &IncomeLettersModule{
		repo: storage.NewIncomSqliteRepository(),
	}
}
func (module *IncomeLettersModule) GetAll(w http.ResponseWriter, r *http.Request) {
	letters, err := module.repo.GetAll()
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

func (module *IncomeLettersModule) Create(w http.ResponseWriter, r *http.Request) {
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
	letterId, err := module.repo.Create(u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"letterId": letterId,
	})
}
