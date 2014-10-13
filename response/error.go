package response

import (
	encodingJSON "encoding/json"
	"errors"
	"fmt"
	"net/http"

	netheadErrors "github.com/zimmski/nethead/errors"
)

type erro struct {
	err error
}

var _ Responder = (*erro)(nil)

func NewError(err error) *erro {
	return &erro{
		err: err,
	}
}

func NewErrorString(s string) *erro {
	return &erro{
		err: errors.New(s),
	}
}

func (r *erro) response() (int, map[string]string) {
	status := http.StatusInternalServerError
	body := map[string]string{
		"error":   "internal",
		"message": r.err.Error(),
	}

	switch err := r.err.(type) {
	case *netheadErrors.Error:
		switch err.Type {
		case netheadErrors.NotFound:
			status = http.StatusNotFound
			body["error"] = err.Type.String()
			body["message"] = err.Message
		}
	}

	return status, body
}

func (r *erro) MarshalJSON() ([]byte, error) {
	_, body := r.response()

	return []byte(fmt.Sprintf(`{"error":%q,"message":%q}`, body["error"], body["message"])), nil
}

func (r *erro) Respond(w http.ResponseWriter) {
	status, body := r.response()

	ret, err := encodingJSON.Marshal(body)
	if err != nil {
		ret = []byte(err.Error())
	}

	w.WriteHeader(status)
	w.Write(ret)
}

type erros struct {
	errs []*netheadErrors.Error
}

var _ Responder = (*erros)(nil)

func Newerros(errs []*netheadErrors.Error) *erros {
	return &erros{
		errs: errs,
	}
}

func (r *erros) Respond(w http.ResponseWriter) {
	status := http.StatusBadRequest
	body := map[string]interface{}{
		"erros": r.errs,
	}

	ret, err := encodingJSON.Marshal(body)
	if err != nil {
		ret = []byte(err.Error())
	}

	w.WriteHeader(status)
	w.Write(ret)
}
