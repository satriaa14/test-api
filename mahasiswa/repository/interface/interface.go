package _interface

import (
	"io"

	"github.com/satriaa14/test-api/mahasiswa/model"
)

// ReadWriter is an interface having the set of methods, which are defined inside the sqlite.
type ReadWriter interface {
	io.Closer
	CreateMahasiswa(req model.Mahasiswa) error
	GetMahasiswa() ([]model.Mahasiswa, error)
	GetMahasiswaByID(req string) (model.Mahasiswa, error)
	DeleteMahasiswa(req string) error
	UpdateMahasiswa() error
}
