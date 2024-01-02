package user_api

import (
	"fmt"
	"net/http"
	"text/template"

	"mehrangcode.ir/office/internal/storage"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/utils"
)

type Page struct {
	Users []types.UserViewModel
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	repo := storage.NewUserSqliteRepo()
	Users, err := repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	templatePath, err := template.ParseFiles("./internal/templates/user.html")
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(templatePath, err)
	tmpl.Execute(w, Users)
	// utils.WriteJson(w, 200, Users)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user types.UserDTO = types.UserDTO{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	repo := storage.NewUserSqliteRepo()
	userId, err := repo.Create(user)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	htmlStr := fmt.Sprintf("<p>%s) %s - %s</p>", userId, user.Name, user.Email)
	tmpl, _ := template.New("t").Parse(htmlStr)
	tmpl.Execute(w, nil)
}
