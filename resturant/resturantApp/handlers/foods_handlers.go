package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"mehrangcode.ir/resturant/app/models"
	"mehrangcode.ir/resturant/app/repositories"
	"mehrangcode.ir/resturant/app/storage"
	"mehrangcode.ir/resturant/app/utils"
)

type foodsApi struct {
	repo repositories.FoodRepo
}
type FoodHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	ChangeStatus(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewFoodsHandler() FoodHandlers {
	return &foodsApi{repo: storage.NewFoodSqliteDB()}
}

func (h *foodsApi) GetAll(w http.ResponseWriter, r *http.Request) {
	Foods, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Foods)
}

func (h *foodsApi) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "foodId")
	Food, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Food)
}

func (h *foodsApi) Create(w http.ResponseWriter, r *http.Request) {
	var newFood models.FoodDTO
	var err error
	err = utils.ReadJson(w, r, &newFood)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if newFood.Name == "" {
		utils.ResponseToError(w, errors.New("food name is required"), http.StatusBadRequest)
		return
	}
	foodId, err := h.repo.Create(newFood)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"foodId": foodId,
	})
}

func (h *foodsApi) Update(w http.ResponseWriter, r *http.Request) {
	foodId := chi.URLParam(r, "foodId")
	var foodDTO models.FoodDTO
	var err error
	err = utils.ReadJson(w, r, &foodDTO)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if foodId == "" {
		utils.ResponseToError(w, errors.New("food name is required"), http.StatusBadRequest)
		return
	}
	foodDTO.ID = foodId
	err = h.repo.Update(foodId, foodDTO)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "food is uptaed",
	})
}

func (h *foodsApi) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	foodId := chi.URLParam(r, "foodId")
	var err error
	status, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if foodId == "" {
		utils.ResponseToError(w, errors.New("food name is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.ChangeStatus(foodId, status)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "status chnaged",
	})
}
func (h *foodsApi) Delete(w http.ResponseWriter, r *http.Request) {
	foodId := chi.URLParam(r, "foodId")
	err := h.repo.Delete(foodId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "food Deleted",
	})
}
