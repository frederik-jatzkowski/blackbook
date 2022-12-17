package util

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/frederik-jatzkowski/blackbook/database"
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

func WriteResponse(w http.ResponseWriter, user *database.User, errs ...string) {
	type ResponseUser struct {
		ID        uint   `json:"id"`
		Active    bool   `json:"active"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}
	type ResponseBody struct {
		Ok     bool          `json:"ok"`
		Errors []string      `json:"errors"`
		User   *ResponseUser `json:"user"`
	}

	body := ResponseBody{
		Ok:     len(errs) == 0,
		Errors: errs,
	}

	if user != nil {
		body.User = &ResponseUser{
			ID:        user.ID,
			Active:    user.Active,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
		}
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
