package sqlite3

import (
	"fmt"
	"strings"

	"github.com/satriaa14/test-api/mahasiswa/model"
)

func (rw *sqLiteReadWriter) CreateMahasiswa(req model.Mahasiswa) error {

	_, err := rw.sqLite.Exec(insertMahasiswa, req.NIM, req.Name, req.Class, req.Phone, req.CreatedAt, req.CreatedBy, req.UpdatedAt, req.UpdatedBy)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed:") {
			return fmt.Errorf("NIM %s is already exist, try another NIM", req.NIM)
		}
		return err
	}

	return nil
}
