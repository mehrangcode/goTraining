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

type coursesApi struct {
	repo repositories.CourseRepo
}

type CourseHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewCourseHandler() CourseHandler {
	return &coursesApi{repo: storage.NewCourseSqliteDB()}
}

func (h *coursesApi) Create(w http.ResponseWriter, r *http.Request) {
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&course); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, course)
}

func (h *coursesApi) GetAll(w http.ResponseWriter, r *http.Request) {
	courses, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, courses)
}

func (h *coursesApi) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "courseID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	course, err := h.repo.GetByID(uint(id))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJson(w, http.StatusOK, course)
}

func (h *coursesApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "courseID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	var course models.Course
	if err := json.NewDecoder(r.Body).Decode(&course); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	course.ID = uint(id)
	if err := h.repo.Update(&course); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, course)
}

func (h *coursesApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "courseID"))
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
