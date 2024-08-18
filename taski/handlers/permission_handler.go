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

type permissionsApi struct {
	repo repositories.PermissionRepo
}

type PermissionHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewPermissionHandler() PermissionHandlers {
	return &permissionsApi{repo: storage.NewPermissionSqliteDB()}
}

func (h *permissionsApi) Create(w http.ResponseWriter, r *http.Request) {
	var permission models.Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&permission); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, permission)
}

func (h *permissionsApi) GetAll(w http.ResponseWriter, r *http.Request) {
	permissions, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, permissions)
}

func (h *permissionsApi) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "permissionID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	permission, err := h.repo.GetByID(uint(id))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJson(w, http.StatusOK, permission)
}

func (h *permissionsApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "permissionID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	var permission models.Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	permission.ID = uint(id)
	if err := h.repo.Update(&permission); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, permission)
}

func (h *permissionsApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "permissionID"))
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
