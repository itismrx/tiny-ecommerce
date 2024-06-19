package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itismrx/tiny-ecommerce/types"
	"github.com/itismrx/tiny-ecommerce/utils"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// get json payload
	var payload types.RegisterdUserPaylodad
	err := utils.ParseJson(r, payload)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// check if the user exist

	// if it doesn't; create the new user
}
