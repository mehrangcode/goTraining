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

type food_categoriesApi struct {
	repo repositories.FoodCategoriesRepo
}
type FoodCategoryHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	ChangeStatus(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewFoodCategoriesHandler() FoodCategoryHandlers {
	return &food_categoriesApi{repo: storage.NewFoodCategorySqliteDB()}
}

func (h *food_categoriesApi) GetAll(w http.ResponseWriter, r *http.Request) {
	List, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, List)
}

func (h *food_categoriesApi) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "catId")
	Item, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Item)
}

func (h *food_categoriesApi) Create(w http.ResponseWriter, r *http.Request) {
	var newItem models.FoodCategoryDTO
	var err error
	err = utils.ReadJson(w, r, &newItem)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if newItem.Title == "" {
		utils.ResponseToError(w, errors.New("title is required"), http.StatusBadRequest)
		return
	}
	catId, err := h.repo.Create(newItem)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"catId": catId,
	})
}

func (h *food_categoriesApi) Update(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "catId")
	var dto models.FoodCategoryDTO
	var err error
	err = utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if catId == "" {
		utils.ResponseToError(w, errors.New("food name is required"), http.StatusBadRequest)
		return
	}
	dto.ID = catId
	err = h.repo.Update(catId, dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, models.FoodCategoryViewModel(dto))
}

func (h *food_categoriesApi) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "catId")
	var err error
	status, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if catId == "" {
		utils.ResponseToError(w, errors.New("food name is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.ChangeStatus(catId, status)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "status chnaged",
	})
}
func (h *food_categoriesApi) Delete(w http.ResponseWriter, r *http.Request) {
	catId := chi.URLParam(r, "catId")
	err := h.repo.Delete(catId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "food Deleted",
	})
}
