package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
	"mehrang.ir/taski/storage"
	"mehrang.ir/taski/utils"
)

type usersApi struct {
	repo repositories.UserRepoInterface
}

type UserHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewUsersHandler() UserHandlers {
	return &usersApi{repo: storage.NewUserSqliteDB()}
}

func (h *usersApi) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	for i := range users {
		users[i].FullName = users[i].FirstName + " " + users[i].LastName
	}
	utils.WriteJson(w, http.StatusOK, users)
}

func (h *usersApi) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	user, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	user.FullName = user.FirstName + " " + user.LastName
	utils.WriteJson(w, http.StatusOK, user)
}

func (h *usersApi) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if !ValidateBody(user) {
		utils.ResponseToError(w, http.ErrMissingFile, http.StatusBadRequest)
		return
	}
	user.Password = string(hashedPassword)
	if err := h.repo.Create(&user); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, user)
}

func (h *usersApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if !ValidateBody(user) {
		utils.ResponseToError(w, http.ErrMissingFile, http.StatusBadRequest)
		return
	}
	user.ID = uint(id)
	if err := h.repo.Update(&user); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, user)
}

func (h *usersApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(id); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusNoContent, nil)
}

func ValidateBody(data models.User) bool {

	if data.FirstName == "" || data.LastName == "" || data.Phone == "" {
		return false
	}

	return true

}
