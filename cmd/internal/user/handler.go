package user

import (
	"net/http"

	"app.go/cmd/internal/handlers"
	"app.go/cmd/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

const (
	usersURL ="/users"
	userURL = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler{
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(usersURL, h.UpdateUser)
	router.PATCH(usersURL, h.PartiallyUpdateUser)
	router.DELETE(usersURL, h.DeleteUser)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("list users"))
}

func (h *handler) GetUserByUUID (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("user by uuid"))
}
func (h *handler) CreateUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("create user"))
}
func (h *handler) UpdateUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("update user"))
}
func (h *handler) PartiallyUpdateUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("partially update user"))
}
func (h *handler) DeleteUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("delete user"))
}
