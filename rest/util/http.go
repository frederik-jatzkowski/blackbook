package util

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody[T interface{}](w http.ResponseWriter, r *http.Request, method string) (T, error) {
	var (
		result T
		data   []byte
		err    error
	)

	// check method
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return result, err
	}

	// read body
	data, err = io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return result, err
	}

	// parse body
	err = json.Unmarshal(data, &result)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return result, err
	}

	return result, err
}

type ResponseBody[T interface{}] struct {
	Ok      bool     `json:"ok"`
	Success *string  `json:"success"`
	Errors  []string `json:"errors"`
	Payload T        `json:"payload"`
}

func WriteResponse[T interface{}](w http.ResponseWriter, payload T, success *string, errs ...string) {

	body := ResponseBody[T]{
		Ok:      len(errs) == 0,
		Success: success,
		Payload: payload,
		Errors:  errs,
	}

	data, _ := json.Marshal(body)

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func MustMarshal(w http.ResponseWriter, message interface{}) {
	data, _ := json.Marshal(message)

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
