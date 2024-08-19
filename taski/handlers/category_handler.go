package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"mehrang.ir/taski/models"
	"mehrang.ir/taski/repositories"
	"mehrang.ir/taski/storage"
	"mehrang.ir/taski/utils"
)

type categoriesApi struct {
	repo repositories.CategoryRepo
}

type CategoryHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCategoryHandler() CategoryHandlers {
	return &categoriesApi{repo: storage.NewCategorySqliteDB()}
}

func (h *categoriesApi) Create(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&category); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, category)
}

func (h *categoriesApi) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, categories)
}

func (h *categoriesApi) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "categoryID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	category, err := h.repo.GetByID(uint(id))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJson(w, http.StatusOK, category)
}

func (h *categoriesApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "categoryID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	var category models.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	category.ID = uint(id)
	if err := h.repo.Update(&category); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, category)
}

func (h *categoriesApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "categoryID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Delete(uint(id)); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
