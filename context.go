package nethead

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/zimmski/nethead/response"
)

type Context struct {
	controllers map[string]Controller
	models      map[string]Model
	router      *mux.Router
}

func New(router *mux.Router) *Context {
	return &Context{
		controllers: make(map[string]Controller),
		models:      make(map[string]Model),
		router:      router,
	}
}

func handleResponse(ctx *Context, r func(ctx *Context, req *http.Request) response.Responder) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		r(ctx, req).Respond(w)
	}
}

func (ctx *Context) addControllerToRoute(contr Controller) {
	uid := contr.UID()

	uri := fmt.Sprintf("/%s", uid)
	idURI := fmt.Sprintf("%s/{%sID:[0-9]+}", uri, uid)

	r := ctx.router

	if c, ok := contr.(AllController); ok {
		r.HandleFunc(uri, handleResponse(ctx, c.All)).Methods("GET")
	}
	if c, ok := contr.(CreateController); ok {
		r.HandleFunc(uri, handleResponse(ctx, c.Create)).Methods("POST")
		r.HandleFunc(uri+"/new", handleResponse(ctx, c.CreateForm)).Methods("GET")
	}
	if c, ok := contr.(DeleteController); ok {
		r.HandleFunc(idURI, handleResponse(ctx, c.Delete)).Methods("DELETE")
	}
	if c, ok := contr.(EditController); ok {
		r.HandleFunc(idURI, handleResponse(ctx, c.Edit)).Methods("PUT")
		r.HandleFunc(idURI+"edit", handleResponse(ctx, c.EditForm)).Methods("GET")
	}
	if c, ok := contr.(OneController); ok {
		r.HandleFunc(idURI, handleResponse(ctx, c.One)).Methods("GET")
	}
}

func (ctx *Context) Controller(uid string) Controller {
	return ctx.controllers[uid]
}

func (ctx *Context) Model(uid string) Model {
	return ctx.models[uid]
}

func (ctx *Context) Router() *mux.Router {
	return ctx.router
}

func (ctx *Context) RegisterController(contr Controller) error {
	uid := contr.UID()

	if _, ok := ctx.controllers[uid]; ok {
		return errors.New("already registered")
	}

	ctx.controllers[uid] = contr

	ctx.addControllerToRoute(contr)

	return nil
}

func (ctx *Context) RegisterModel(m Model) error {
	uid := m.UID()

	if _, ok := ctx.models[uid]; ok {
		return errors.New("already registered")
	}

	ctx.models[uid] = m

	return nil
}
