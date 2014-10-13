package nethead

import (
	"github.com/zimmski/nethead/response"
	"net/http"
)

type Controller interface {
	UID() string
}

type AllController interface {
	All(ctx *Context, req *http.Request) response.Responder
}

type CreateController interface {
	Create(ctx *Context, req *http.Request) response.Responder
	CreateForm(ctx *Context, req *http.Request) response.Responder
}

type DeleteController interface {
	Delete(ctx *Context, req *http.Request) response.Responder
}

type EditController interface {
	Edit(ctx *Context, req *http.Request) response.Responder
	EditForm(ctx *Context, req *http.Request) response.Responder
}

type OneController interface {
	One(ctx *Context, req *http.Request) response.Responder
}
