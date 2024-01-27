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

type reservationsApi struct {
	repo repositories.ReservationRepo
}
type ReservationHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	ChangeStatus(w http.ResponseWriter, r *http.Request)
}

func NewReservationsHandler() ReservationHandlers {
	return &reservationsApi{repo: storage.NewReservationSqliteDB()}
}

func (h *reservationsApi) GetAll(w http.ResponseWriter, r *http.Request) {
	Reservations, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Reservations)
}
func (h *reservationsApi) GetByUserId(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "userId")
	Reservations, err := h.repo.GetByUserId(userId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Reservations)
}

func (h *reservationsApi) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "reservationId")
	Reservation, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Reservation)
}

func (h *reservationsApi) Create(w http.ResponseWriter, r *http.Request) {
	var newReservation models.ReservationDTO
	err := utils.ReadJson(w, r, &newReservation)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if newReservation.UserID == "" || newReservation.Guests == 0 || newReservation.Date == "" {
		utils.ResponseToError(w, errors.New("reservation fields is required"), http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(newReservation)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"reservationId": id,
	})
}

func (h *reservationsApi) Update(w http.ResponseWriter, r *http.Request) {
	reservationId := chi.URLParam(r, "reservationId")
	var dto models.ReservationDTO
	err := utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if reservationId == "" || dto.Guests == 0 || dto.Date == "" {
		utils.ResponseToError(w, errors.New("reservation fields is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(reservationId, dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "reservation successfuly updated",
	})
}

func (h *reservationsApi) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	reservationId := chi.URLParam(r, "reservationId")
	status, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	err = h.repo.ChangeStatus(reservationId, uint(status))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "reservation status updated",
	})
}

func (h *reservationsApi) Delete(w http.ResponseWriter, r *http.Request) {
	reservationId := chi.URLParam(r, "reservationId")
	err := h.repo.Delete(reservationId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "reservation Deleted",
	})
}
