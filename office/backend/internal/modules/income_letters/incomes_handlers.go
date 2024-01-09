package income_letters

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
	IncomeLetters, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, IncomeLetters)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var dto types.IncomeLetterDTO
	var err error
	err = utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if dto.Number <= 0 {
		utils.ResponseToError(w, errors.New("number is required"), http.StatusBadRequest)
		return
	}
	letterId, err := h.repo.Create(dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"letterId": letterId,
	})
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	letterId := chi.URLParam(r, "letterId")
	var dto types.IncomeLetterDTO
	var err error
	err = utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if letterId == "" || dto.Number <= 0 {
		utils.ResponseToError(w, errors.New("label is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(letterId, dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "letteris Updated",
	})
}
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	letterId := chi.URLParam(r, "letterId")
	err := h.repo.Delete(letterId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "Letter Deleted",
	})
}
