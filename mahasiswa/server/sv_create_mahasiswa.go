package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/satriaa14/test-api/mahasiswa/middleware"
	"github.com/satriaa14/test-api/mahasiswa/model"
	"github.com/satriaa14/test-api/mahasiswa/util/logging"
	"github.com/satriaa14/test-api/mahasiswa/util/pattern"
	"github.com/satriaa14/test-api/mahasiswa/util/setup"
)

func (rw *Service) CreateMahasiswa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		setup.SetupCorsResponse(&w, r)

		if pattern.DataURL(r.URL.Path)[4] != create {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// Middle ware
		err := middleware.MiddlewareJwtAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if r.Method != http.MethodPost {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}

		// TODO

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

		request.CreatedAt = time.Now()
		request.CreatedBy = "User"
		request.UpdatedAt = time.Now()
		request.UpdatedBy = "User"

		err = rw.repo.SQLiteReadWriter.CreateMahasiswa(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		w.Write([]byte("Create Success"))
	}
}
