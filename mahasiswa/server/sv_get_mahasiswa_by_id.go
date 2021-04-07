package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/satriaa14/test-api/mahasiswa/middleware"
	"github.com/satriaa14/test-api/mahasiswa/model"
	"github.com/satriaa14/test-api/mahasiswa/util/logging"
	"github.com/satriaa14/test-api/mahasiswa/util/pattern"
	"github.com/satriaa14/test-api/mahasiswa/util/setup"
)

func (rw *Service) GetMahasiswaByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		setup.SetupCorsResponse(&w, r)

		if pattern.DataURL(r.URL.Path)[4] != getall {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// Middle ware
		err := middleware.MiddlewareJwtAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if r.Method != http.MethodGet {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var request model.Mahasiswa
		if err = json.Unmarshal(body, &request); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if request.NIM == "" || len(request.NIM) == 0 {
			http.Error(w, "Please Fill NIM", http.StatusBadRequest)
			return
		}

		// TODO
		resp, err := rw.repo.SQLiteReadWriter.GetMahasiswaByID(request.NIM)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		data, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		w.Write(data)
	}
}
