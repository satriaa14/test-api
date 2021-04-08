package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	apigetall  string = api + mahasiswa + root + getall
	apicreate  string = api + mahasiswa + root + create
	apigetbyid string = api + mahasiswa + root + get + root
	apidelete  string = api + mahasiswa + root + delete + root
	apiupdate  string = api + mahasiswa + root + update + root
)

type Service struct {
	repo repository.Repo
}

func NewMahasiswaService(repo repository.Repo) service.Operate {
	return &Service{repo: repo}
}

func Run() {

	godotenv.Load()
	port := os.Getenv("PORT")

	newRepo, err := repository.NewMahasiswaRepo()
	if err != nil {
		panic(err)
	}

	mahasiswaService := NewMahasiswaService(*newRepo)

	fmt.Println(port)
	if port == "" {
		port = "3000"
	}

	// Handlers
	http.HandleFunc("/", mahasiswaService.Alive())
	http.HandleFunc("/login", mahasiswaService.Login())
	http.HandleFunc(apigetall, mahasiswaService.GetMahasiswa())
	http.HandleFunc(apidelete, mahasiswaService.DeleteMahasiswa())
	http.HandleFunc(apicreate, mahasiswaService.CreateMahasiswa())
	http.HandleFunc(apiupdate, mahasiswaService.UpdateMahasiswa())
	http.HandleFunc(apigetbyid, mahasiswaService.GetMahasiswaByID())

	log.Println("Start server")
	http.ListenAndServe(":"+port, nil)
}
