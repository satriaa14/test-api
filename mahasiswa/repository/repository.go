package repository

import (
	"io"

	_interface "github.com/satriaa14/test-api/mahasiswa/repository/interface"
	"github.com/satriaa14/test-api/mahasiswa/repository/sqlite3"
)

type Repo struct {
	io.Closer
	SQLiteReadWriter _interface.ReadWriter
}

func NewMahasiswaRepo() (*Repo, error) {
	// Database
	readWriter, err := sqlite3.NewSQLiteReadWriter()
	if err != nil {
		return nil, err
	}

	return &Repo{
		SQLiteReadWriter: readWriter,
	}, nil
}

func (r *Repo) Close() {
	r.SQLiteReadWriter.Close()
}
