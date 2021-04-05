package sqlite3

import (
	"database/sql"

	"github.com/satriaa14/test-api/mahasiswa/model"
)

func (rw *sqLiteReadWriter) GetMahasiswa() ([]model.Mahasiswa, error) {
	var resp []model.Mahasiswa

	rows, err := rw.sqLite.Query(getAllMahasiswa)
	if err != nil {
		return resp, err
	}

	for rows.Next() {
		data := model.Mahasiswa{}
		err = rows.Scan(
			&data.NIM,
			&data.Name,
			&data.Class,
			&data.Phone,
			&data.CreatedAt,
			&data.CreatedBy,
			&data.UpdatedAt,
			&data.UpdatedBy,
		)
		if err != nil && err != sql.ErrNoRows {
			return resp, err
		}
		resp = append(resp, data)
	}

	rows.Close()

	return resp, nil
}
