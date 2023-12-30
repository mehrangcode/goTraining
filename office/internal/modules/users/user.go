package users

import (
	"net/http"

	"mehrangcode.ir/office/utils"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	repo := NewSqliteRepo()
	users, err := repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, users)
}
