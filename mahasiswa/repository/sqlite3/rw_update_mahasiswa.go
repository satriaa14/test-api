package sqlite3

import (
	"errors"
	"fmt"
	"strings"

	"github.com/satriaa14/test-api/mahasiswa/model"
)

func (rw *sqLiteReadWriter) UpdateMahasiswa(req model.Mahasiswa) error {

	args, query := rw.updateMahasiswaQBuilder(req, updateMahasiswa)

	args = append(args, req.NIM)

	result, err := rw.sqLite.Exec(query, args...)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("No rows updated")
	}
	return nil
}

func (rw *sqLiteReadWriter) updateMahasiswaQBuilder(req model.Mahasiswa, query string) ([]interface{}, string) {

	const questionMark string = " = ?"

	var fields []string

	var args []interface{}

	var emptyMahasiswa model.Mahasiswa

	if req.Name != emptyMahasiswa.Name {
		fields = append(fields, "name"+questionMark)
		args = append(args, req.Name)
	}

	if req.Class != emptyMahasiswa.Class {
		fields = append(fields, "class"+questionMark)
		args = append(args, req.Class)
	}

	if req.Phone != emptyMahasiswa.Phone {
		fields = append(fields, "phone"+questionMark)
		args = append(args, req.Phone)
	}

	if req.UpdatedAt != emptyMahasiswa.UpdatedAt {
		fields = append(fields, "updated_at"+questionMark)
		args = append(args, req.UpdatedAt)
	}

	if req.UpdatedBy != emptyMahasiswa.UpdatedBy {
		fields = append(fields, "updated_by"+questionMark)
		args = append(args, req.UpdatedBy)
	}

	query = fmt.Sprintf(query, strings.Join(fields, ","))
	return args, query
}
