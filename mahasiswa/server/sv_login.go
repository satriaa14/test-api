package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/satriaa14/test-api/mahasiswa/operation"
	"github.com/satriaa14/test-api/mahasiswa/util/logging"
	"github.com/satriaa14/test-api/mahasiswa/util/pattern"
	"github.com/satriaa14/test-api/mahasiswa/util/setup"
)

func (rw *Service) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		setup.SetupCorsResponse(&w, r)
		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		if pattern.DataURL(r.URL.Path)[1] != login {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var data operation.LoginRequest
		if err = json.Unmarshal(body, &data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := operation.Authenticate(data.UserName, data.Password)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusBadRequest)
			return
		}

		if token == nil {
			http.Error(w, "Token is nil", http.StatusBadRequest)
			return
		}

		w.Write([]byte(token))
	}
}
