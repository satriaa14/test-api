package sqlite3

import (
	"github.com/satriaa14/test-api/mahasiswa/model"
)

func (rw *sqLiteReadWriter) UpdateMahasiswa(req model.Mahasiswa) error {
	_, err := rw.sqLite.Exec(insertMahasiswa, req.NIM, req.Name, req.Class, req.Phone, req.CreatedAt, req.CreatedBy, req.UpdatedAt, req.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}
