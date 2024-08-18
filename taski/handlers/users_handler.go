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
	repo repositories.UserRepo
}

type UserHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	AddRolesToUser(w http.ResponseWriter, r *http.Request)
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
	var updatedUsers []models.UserVieModel
	for _, user := range users {
		uu := models.UserVieModel{
			ID:       user.ID,
			FullName: user.FirstName + " " + user.LastName,
			Phone:    user.Phone,
		}
		updatedUsers = append(updatedUsers, uu)
	}
	utils.WriteJson(w, http.StatusOK, updatedUsers)
}

func (h *usersApi) GetById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	user, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	updatedUser := models.UserVieModel{
		ID:       user.ID,
		FullName: user.FirstName + " " + user.LastName,
		Phone:    user.Phone,
		Roles:    user.Roles,
	}
	utils.WriteJson(w, http.StatusOK, updatedUser)
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
	utils.WriteJson(w, http.StatusCreated, map[string]uint{
		"userID": user.ID,
	})
}

func (h *usersApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))
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
	utils.WriteJson(w, http.StatusOK, map[string]uint{
		"userID": user.ID,
	})
}

func (h *usersApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "userID"))
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

func (h *usersApi) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByPhone(creds.Phone)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		utils.ResponseToError(w, err, http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{"token": token})
}

func (h *usersApi) AddRolesToUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	var roleIDs []uint
	if err := json.NewDecoder(r.Body).Decode(&roleIDs); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	var user models.User
	if user, err = h.repo.AddRolesToUser(uint(userID), roleIDs); err != nil {
		utils.ResponseToError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteJson(w, http.StatusOK, user)
}

// filling user view model object
// func fillUserViewModel(user *models.User, userViewModel *models.UserVieModel) {
// 	userVal := reflect.ValueOf(user).Elem()
// 	userViewModelVal := reflect.ValueOf(userViewModel).Elem()

// 	for i := 0; i < userViewModelVal.NumField(); i++ {
// 		field := userViewModelVal.Type().Field(i)
// 		userField := userVal.FieldByName(field.Name)
// 		if userField.IsValid() {
// 			userViewModelVal.Field(i).Set(userField)
// 		}
// 	}

// 	// Special handling for FullName
// 	fullName := user.FirstName + " " + user.LastName
// 	userViewModelVal.FieldByName("FullName").SetString(fullName)
// }
