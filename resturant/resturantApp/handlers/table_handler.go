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

type tablesApi struct {
	repo repositories.TableRepo
}
type TableHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	ChangeStatus(w http.ResponseWriter, r *http.Request)
	Reservation(w http.ResponseWriter, r *http.Request)
}

func NewTablesHandler() TableHandlers {
	return &tablesApi{repo: storage.NewTableSqliteDB()}
}

func (h *tablesApi) GetAll(w http.ResponseWriter, r *http.Request) {
	Tables, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Tables)
}

func (h *tablesApi) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "tableId")
	Table, err := h.repo.GetById(id)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, Table)
}

func (h *tablesApi) Create(w http.ResponseWriter, r *http.Request) {
	var newTable models.TableDTO
	err := utils.ReadJson(w, r, &newTable)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if newTable.Name == "" || newTable.Capacity == 0 {
		utils.ResponseToError(w, errors.New("table fields is required"), http.StatusBadRequest)
		return
	}
	id, err := h.repo.Create(newTable)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"tableId": id,
	})
}

func (h *tablesApi) Update(w http.ResponseWriter, r *http.Request) {
	tableId := chi.URLParam(r, "tableId")
	var dto models.TableDTO
	err := utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if tableId == "" || dto.Name == "" || dto.Capacity == 0 {
		utils.ResponseToError(w, errors.New("table fields is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Update(tableId, dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	dto.ID = tableId
	utils.WriteJson(w, http.StatusOK, models.TableViewModel(dto))
}

func (h *tablesApi) ChangeStatus(w http.ResponseWriter, r *http.Request) {
	tableId := chi.URLParam(r, "tableId")
	status, err := strconv.Atoi(chi.URLParam(r, "status"))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}
	err = h.repo.ChangeStatus(tableId, uint(status))
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "table status updated",
	})
}
func (h *tablesApi) Reservation(w http.ResponseWriter, r *http.Request) {
	tableId := chi.URLParam(r, "tableId")
	var dto models.ReservationDTO
	err := utils.ReadJson(w, r, &dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
	}
	if tableId == "" || dto.UserID == "" || dto.TableID == "" || dto.Date == "" {
		utils.ResponseToError(w, errors.New("table fields is required"), http.StatusBadRequest)
		return
	}
	err = h.repo.Reservation(tableId, dto)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusBadRequest)
		return
	}

	utils.WriteJson(w, http.StatusOK, map[string]string{
		"msg": "table reservation was successful",
	})
}

func (h *tablesApi) Delete(w http.ResponseWriter, r *http.Request) {
	tableId := chi.URLParam(r, "tableId")
	err := h.repo.Delete(tableId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "table Deleted",
	})
}
