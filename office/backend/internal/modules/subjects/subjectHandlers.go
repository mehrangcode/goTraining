package subjects

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mehrangcode.ir/office/internal/types"
	"mehrangcode.ir/office/utils"
)

type Handler struct {
	repo Repository
}

func NewHandlers(repo Repository) *Handler {
	return &Handler{
		repo: repo,
	}
}
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	Subjects, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Subjects)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto types.SubjectDTO
	var err error
	err = utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if dto.Label == "" {
		utils.ResponseToError(w, errors.New("subject fields is required"), http.StatusBadRequest)
		return
	}
	itemId, err := h.repo.Create(dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"itemId": itemId,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	subjectId := chi.URLParam(r, "subjectId")
	var dto types.SubjectDTO
	var err error
	err = utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if subjectId == "" || dto.Label == "" {
		utils.ResponseToError(w, errors.New("label is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(subjectId, dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, types.SubjectViewModel{
		ID:    subjectId,
		Label: dto.Label,
	})
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	subjectId := chi.URLParam(r, "subjectId")
	err := h.repo.Delete(subjectId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "subject Deleted",
	})
}
