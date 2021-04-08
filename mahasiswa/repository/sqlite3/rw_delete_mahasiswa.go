package sqlite3

import "errors"

func (rw *sqLiteReadWriter) DeleteMahasiswa(req string) error {

	result, err := rw.sqLite.Exec(deleteMahasiswaByID, req)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("No rows deleted")
	}

	return nil
}
