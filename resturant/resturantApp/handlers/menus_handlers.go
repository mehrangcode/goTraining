package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"mehrangcode.ir/resturant/app/models"
	"mehrangcode.ir/resturant/app/repositories"
	"mehrangcode.ir/resturant/app/storage"
	"mehrangcode.ir/resturant/app/utils"
)

type menusApi struct {
	repo repositories.MenuRepo
}
type MenuHandlers interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	// GetById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	// Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func NewMenusHandler() MenuHandlers {
	return &menusApi{repo: storage.NewMenuSqliteDB()}
}

func (h *menusApi) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.GetAll()
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, 200, list)
}

// func (h *menusApi) GetById(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "menuId")
// 	Menu, err := h.repo.GetById(id)
// 	if err != nil {
// 		utils.ResponseToError(w, err, http.StatusInternalServerError)
// 		return
// 	}
// 	utils.WriteJson(w, 200, Menu)
// }

func (h *menusApi) Create(w http.ResponseWriter, r *http.Request) {
	var newMenu models.MenuDTO
	var err error
	err = utils.ReadJson(w, r, &newMenu)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	if newMenu.Title == "" {
		utils.ResponseToError(w, errors.New("menu fields is required"), http.StatusBadRequest)
		return
	}
	menuId, err := h.repo.Create(newMenu)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, map[string]string{
		"menuId": menuId,
	})
}

//	func (h *menusApi) Update(w http.ResponseWriter, r *http.Request) {
//		menuId := chi.URLParam(r, "menuId")
//		var menuDTO models.MenuDTO
//		var err error
//		err = utils.ReadJson(w, r, &menuDTO)
//		if err != nil {
//			utils.ResponseToError(w, err, http.StatusBadRequest)
//		}
//		if menuId == "" || menuDTO.Title == "" {
//			utils.ResponseToError(w, errors.New("menu fields is required"), http.StatusBadRequest)
//			return
//		}
//		err = h.repo.Update(menuId, menuDTO)
//		if err != nil {
//			utils.ResponseToError(w, err, http.StatusBadRequest)
//			return
//		}
//		menuDTO.ID = menuId
//		utils.WriteJson(w, http.StatusOK, map[string]string{
//			"msg": "menu get updated",
//		})
//	}
func (h *menusApi) Delete(w http.ResponseWriter, r *http.Request) {
	menuId := chi.URLParam(r, "menuId")
	err := h.repo.Delete(menuId)
	if err != nil {
		utils.ResponseToError(w, err, http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]string{
		"message": "menu Deleted",
	})
}
