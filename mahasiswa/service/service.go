package service

import (
	"net/http"
)

type Operate interface {
	PreService
	MahasiswaService
}

type PreService interface {
	Alive() http.HandlerFunc
	Login() http.HandlerFunc
}

type MahasiswaService interface {
	GetMahasiswa() http.HandlerFunc
	DeleteMahasiswa() http.HandlerFunc
	CreateMahasiswa() http.HandlerFunc
	UpdateMahasiswa() http.HandlerFunc
	GetMahasiswaByID() http.HandlerFunc
}
