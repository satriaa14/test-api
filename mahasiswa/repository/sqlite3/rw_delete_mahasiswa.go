package sqlite3

import "database/sql"

func (rw *sqLiteReadWriter) DeleteMahasiswa(req string) error {

	_, err := rw.sqLite.Exec(getMahasiswaByID, req)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
