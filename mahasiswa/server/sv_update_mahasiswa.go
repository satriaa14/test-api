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

func (rw *Service) UpdateMahasiswa() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var id string

		var pattern = pattern.DataURL(r.URL.Path)

		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		setup.SetupCorsResponse(&w, r)

		if pattern[4] != update {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		// Middle ware
		err := middleware.MiddlewareJwtAuth(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		if r.Method != http.MethodPut {
			http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
			return
		}
		if len(pattern) < 6 {
			http.Error(w, "URL Not valid", http.StatusNotFound)
			return
		}

		if pattern[5] == "" || len(pattern) == 0 {
			http.Error(w, "Please Fill NIM", http.StatusBadRequest)
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

		id = pattern[5]

		request.NIM = id
		request.UpdatedAt = time.Now()
		request.UpdatedBy = "Updated User"

		err = rw.repo.SQLiteReadWriter.UpdateMahasiswa(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		w.Write([]byte("Create Success"))
	}
}
