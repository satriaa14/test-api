package sqlite3

import (
	"database/sql"

	"github.com/satriaa14/test-api/mahasiswa/model"
)

func (rw *sqLiteReadWriter) GetMahasiswaByID(req string) (model.Mahasiswa, error) {
	var resp model.Mahasiswa

	err := rw.sqLite.QueryRow(getMahasiswaByID, req).Scan(
		&resp.NIM,
		&resp.Name,
		&resp.Class,
		&resp.Phone,
		&resp.CreatedAt,
		&resp.CreatedBy,
		&resp.UpdatedAt,
		&resp.UpdatedBy,
	)
	if err != nil && err != sql.ErrNoRows {
		return resp, err
	}

	return resp, nil
}
