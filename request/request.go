package request

import (
	"encoding/json"
	"net/http"
)

func DecodeJsonBody(r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)

	err := d.Decode(v)
	if err != nil {
		return err
	}

	r.Body.Close()

	return nil
}
