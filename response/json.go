package response

import (
	encodingJSON "encoding/json"
	"net/http"
)

type json struct {
	data interface{}
}

var _ Responder = (*json)(nil)

func NewJSON(data interface{}) *json {
	return &json{
		data: data,
	}
}

func (r *json) Respond(w http.ResponseWriter) {
	ret, err := encodingJSON.Marshal(map[string]interface{}{
		"success": true,
		"data":    r.data,
	})
	if err != nil {
		ret = []byte(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ret)
}
