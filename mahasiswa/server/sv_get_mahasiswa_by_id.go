package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/satriaa14/test-api/mahasiswa/middleware"
	"github.com/satriaa14/test-api/mahasiswa/util/logging"
	"github.com/satriaa14/test-api/mahasiswa/util/pattern"
	"github.com/satriaa14/test-api/mahasiswa/util/setup"
)

func (rw *Service) GetMahasiswaByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request string

		var pattern = pattern.DataURL(r.URL.Path)

		log.Println(fmt.Sprintf("%+v", logging.LogRequestClient(w, r)))

		setup.SetupCorsResponse(&w, r)

		if pattern[4] != get {
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

		if len(pattern) < 6 {
			http.Error(w, "URL Not valid", http.StatusNotFound)
			return
		}

		if pattern[5] == "" || len(pattern) == 0 {
			http.Error(w, "Please Fill NIM", http.StatusBadRequest)
			return
		}

		request = pattern[5]

		// TODO
		resp, err := rw.repo.SQLiteReadWriter.GetMahasiswaByID(request)
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
