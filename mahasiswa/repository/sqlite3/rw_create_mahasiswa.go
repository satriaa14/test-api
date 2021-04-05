package sqlite3

import (
	"github.com/satriaa14/test-api/mahasiswa/model"
)

func (rw *sqLiteReadWriter) CreateMahasiswa(req model.Mahasiswa) error {

	// returning statement
	stmt, err := rw.sqLite.Prepare(insertMahasiswa)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(req.NIM, req.Name, req.Class, req.Phone, req.CreatedAt, req.CreatedBy, req.UpdatedAt, req.UpdatedBy)
	if err != nil {
		return err
	}

	return nil
}
