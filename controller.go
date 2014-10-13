package nethead

import (
	"net/http"
)

type Controller interface {
	UID() string
}

type AllController interface {
	All(ctx *Context, req *http.Request) interface{}
}

type CreateController interface {
	Create(ctx *Context, req *http.Request) interface{}
	CreateForm(ctx *Context, req *http.Request) interface{}
}

type DeleteController interface {
	Delete(ctx *Context, req *http.Request) interface{}
}

type EditController interface {
	Edit(ctx *Context, req *http.Request) interface{}
	EditForm(ctx *Context, req *http.Request) interface{}
}

type OneController interface {
	One(ctx *Context, req *http.Request) interface{}
}
