package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zimmski/nethead/response"
)

type Controller interface {
	UID() string
}

func handleResponse(r func(req *http.Request) response.Responder) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		r(req).Respond(w)
	}
}

func AddControllerToRoute(controller Controller, router *mux.Router) {
	uid := controller.UID()

	idURI := fmt.Sprintf("/{%sID:[0-9]+}/", uid)

	if c, ok := controller.(AllController); ok {
		router.HandleFunc("/", handleResponse(c.All)).Methods("GET")
	}
	if c, ok := controller.(CreateController); ok {
		router.HandleFunc("/", handleResponse(c.Create)).Methods("POST")
		router.HandleFunc("/new/", handleResponse(c.CreateForm)).Methods("GET")
	}
	if c, ok := controller.(DeleteController); ok {
		router.HandleFunc(idURI, handleResponse(c.Delete)).Methods("DELETE")
	}
	if c, ok := controller.(EditController); ok {
		router.HandleFunc(idURI, handleResponse(c.Edit)).Methods("PUT")
		router.HandleFunc(idURI+"edit/", handleResponse(c.EditForm)).Methods("GET")
	}
	if c, ok := controller.(OneController); ok {
		router.HandleFunc(idURI, handleResponse(c.One)).Methods("GET")
	}
}

type AllController interface {
	All(req *http.Request) response.Responder
}

type CreateController interface {
	Create(req *http.Request) response.Responder
	CreateForm(req *http.Request) response.Responder
}

type DeleteController interface {
	Delete(req *http.Request) response.Responder
}

type EditController interface {
	Edit(req *http.Request) response.Responder
	EditForm(req *http.Request) response.Responder
}

type OneController interface {
	One(req *http.Request) response.Responder
}
