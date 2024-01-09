package users

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/utils"
)

type UserHandler struct {
	repo UserRepository
}

func NewHandler(repo UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}
func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	Users, err := h.repo.GetAll()
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

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	// var user types.UserDTO = types.UserDTO{
	// 	Name:     r.PostFormValue("name"),
	// 	Email:    r.PostFormValue("email"),
	// 	Password: r.PostFormValue("password"),
	// }
	var newUser types.UserDTO
	var err error
	err = utils.ReadJson(w, r, &newUser)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if newUser.Name == "" || newUser.Email == "" || newUser.Password == "" {
		utils.ResponseToError(w, errors.New("user fields is required"), http.StatusBadRequest)
		return
	}
	userId, err := h.repo.Create(newUser)
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

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	var userDTO types.UserDTO
	var err error
	err = utils.ReadJson(w, r, &userDTO)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if userId == "" || userDTO.Name == "" || userDTO.Email == "" {
		utils.ResponseToError(w, errors.New("user fields is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(userId, userDTO)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, types.UserViewModel{
		ID:    userId,
		Name:  userDTO.Name,
		Email: userDTO.Email,
	})
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	err := h.repo.Delete(userId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "user Deleted",
	})
}
