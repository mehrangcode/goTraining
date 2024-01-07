package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	// templatePath, err := template.ParseFiles("./internal/templates/user.html")
	// if err != nil {
	// 	utils.ResponseToError(w, err, http.StatusInternalServerError)
	// 	return
	// }
	// tmpl := template.Must(templatePath, err)
	// tmpl.Execute(w, Users)
	utils.WriteJson(w, 200, Users)
}

func Create(w http.ResponseWriter, r *http.Request) {
	// var user types.UserDTO = types.UserDTO{
	// 	Name:     r.PostFormValue("name"),
	// 	Email:    r.PostFormValue("email"),
	// 	Password: r.PostFormValue("password"),
	// }
	var u types.UserDTO
	var err error
	err = utils.ReadJson(w, r, &u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if u.Name == "" || u.Email == "" || u.Password == "" {
		utils.ResponseToError(w, errors.New("user fields is required"), http.StatusBadRequest)
		return
	}
	repo := storage.NewUserSqliteRepo()
	userId, err := repo.Create(u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	// htmlStr := fmt.Sprintf("<p>%s) %s - %s</p>", userId, user.Name, user.Email)
	// tmpl, _ := template.New("t").Parse(htmlStr)
	// tmpl.Execute(w, nil)
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"userId": userId,
	})
}

func Update(w http.ResponseWriter, r *http.Request) {
	repo := storage.NewUserSqliteRepo()
	userId := chi.URLParam(r, "userId")
	var u types.UserDTO
	var err error
	err = utils.ReadJson(w, r, &u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	u.ID = userId
	if u.ID == "" || u.Name == "" || u.Email == "" {
		utils.ResponseToError(w, errors.New("user fields is required"), http.StatusBadRequest)
		return
	}
	err = repo.Update(u)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]types.UserViewModel{
		"user": types.UserViewModel(u),
	})
}
func Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	repo := storage.NewUserSqliteRepo()
	err := repo.Delete(userId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "user Deleted",
	})
}
