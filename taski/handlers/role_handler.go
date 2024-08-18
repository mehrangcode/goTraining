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

type rolesApi struct {
	repo repositories.RoleRepo
}

type RoleHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	AddPermissionsToRole(w http.ResponseWriter, r *http.Request)
}

func NewRoleHandler() RoleHandlers {
	return &rolesApi{repo: storage.NewRoleSqliteDB()}
}

func (h *rolesApi) Create(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	if err := h.repo.Create(&role); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, role)
}

func (h *rolesApi) GetAll(w http.ResponseWriter, r *http.Request) {
	roles, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, roles)
}

func (h *rolesApi) GetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	role, err := h.repo.GetByID(uint(id))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusNotFound)
		return
	}
	utils.WriteJson(w, http.StatusOK, role)
}

func (h *rolesApi) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	role.ID = uint(id)
	if err := h.repo.Update(&role); err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, role)
}

func (h *rolesApi) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "roleID"))
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

func (h *rolesApi) AddPermissionsToRole(w http.ResponseWriter, r *http.Request) {
	roleID, err := strconv.Atoi(chi.URLParam(r, "roleID"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	var permissionIDs []int
	if err := json.NewDecoder(r.Body).Decode(&permissionIDs); err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	var role models.Role
	if role, err = h.repo.AddPermissionsToRole(roleID, permissionIDs); err != nil {
		utils.ResponseToError(w, err, http.StatusNotFound)
		return
	}

	utils.WriteJson(w, http.StatusOK, role)
}
