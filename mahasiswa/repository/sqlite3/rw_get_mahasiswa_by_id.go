package sqlite3

import (
	"database/sql"
	"fmt"

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
	if err != nil {
		if err == sql.ErrNoRows {
			return resp, fmt.Errorf("mahasiswa with NIM %s is not exist", req)
		}
		return resp, err
	}

	return resp, nil
}
