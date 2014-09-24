package response

import (
	encodingJSON "encoding/json"
	"fmt"
	"net/http"

	"github.com/zimmski/nethead"
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

func (r *erro) response() (int, map[string]string) {
	status := http.StatusInternalServerError
	body := map[string]string{
		"error":   "internal",
		"message": r.err.Error(),
	}

	switch err := r.err.(type) {
	case *nethead.Error:
		switch err.Type {
		case nethead.NotFound:
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

type errors struct {
	errs []*nethead.Error
}

var _ Responder = (*errors)(nil)

func NewErrors(errs []*nethead.Error) *errors {
	return &errors{
		errs: errs,
	}
}

func (r *errors) Respond(w http.ResponseWriter) {
	status := http.StatusInternalServerError
	body := map[string]interface{}{
		"errors": r.errs,
	}

	ret, err := encodingJSON.Marshal(body)
	if err != nil {
		ret = []byte(err.Error())
	}

	w.WriteHeader(status)
	w.Write(ret)
}
