package server

import (
	"log"
	"net/http"

	"github.com/satriaa14/test-api/mahasiswa/repository"
	"github.com/satriaa14/test-api/mahasiswa/service"
)

const (
	api       string = "/api/v1/"
	root      string = "/"
	login     string = "login"
	mahasiswa string = "mahasiswa"
	getall    string = "all"
	get       string = "id"
	create    string = "create"
	update    string = "update"
	delete    string = "delete"
)

type Service struct {
	repo repository.Repo
}

func NewMahasiswaService(repo repository.Repo) service.Operate {
	return &Service{repo: repo}
}

func Run() {

	newRepo, err := repository.NewMahasiswaRepo()
	if err != nil {
		panic(err)
	}

	mahasiswaService := NewMahasiswaService(*newRepo)

	// Handlers
	http.HandleFunc("/", mahasiswaService.Alive())
	http.HandleFunc("/login", mahasiswaService.Login())
	http.HandleFunc(api+mahasiswa+root+get, mahasiswaService.GetMahasiswaByID())
	http.HandleFunc(api+mahasiswa+root+getall, mahasiswaService.GetMahasiswa())
	http.HandleFunc(api+mahasiswa+root+create, mahasiswaService.CreateMahasiswa())
	// http.HandleFunc(api+mahasiswa+root+update, mahasiswaHandler())
	// http.HandleFunc(api+mahasiswa+root+delete, mahasiswaHandler())

	log.Println("Start server")
	http.ListenAndServe(":9000", nil)
}
