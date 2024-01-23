package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"mehrangcode.ir/resturant/app/models"
	"mehrangcode.ir/resturant/app/repositories"
	"mehrangcode.ir/resturant/app/storage"
	"mehrangcode.ir/resturant/app/utils"
)

type usersApi struct {
	repo repositories.UserRepo
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
	Users, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Users)
}

func (h *usersApi) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "userId")
	User, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, User)
}

func (h *usersApi) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.UserDTO
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
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"userId": userId,
	})
}

func (h *usersApi) Update(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.Atoi(chi.URLParam(r, "userId"))
	var userDTO models.UserDTO
	var err error
	err = utils.ReadJson(w, r, &userDTO)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if userId == 0 || userDTO.Name == "" || userDTO.Email == "" {
		utils.ResponseToError(w, errors.New("user fields is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(fmt.Sprint(userId), userDTO)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, models.UserViewModel{
		ID:    uint(userId),
		Name:  userDTO.Name,
		Email: userDTO.Email,
	})
}
func (h *usersApi) Delete(w http.ResponseWriter, r *http.Request) {
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
