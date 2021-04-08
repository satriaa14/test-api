package sqlite3

import (
	"database/sql"
	"log"

	_interface "github.com/satriaa14/test-api/mahasiswa/repository/interface"

	_ "github.com/mattn/go-sqlite3"
)

// dbReadWriter is a struct having the sql db parameter
type sqLiteReadWriter struct {
	sqLite *sql.DB
}

const (
	// Create table if not exists
	sqlTable = `CREATE TABLE IF NOT EXISTS mahasiswa(
					nim TEXT NOT NULL PRIMARY KEY,
					name TEXT,
					class TEXT,
					phone VARCHAR(14),
					created_at DATETIME,
					created_by TEXT,
					updated_at DATETIME,
					updated_by TEXT
				);`

	// Query operation
	insertMahasiswa = `INSERT INTO mahasiswa(
					nim,
					name,
					class,
					phone,
					created_at,
					created_by,
					updated_at,updated_by
				) VALUES(?,?,?,?,?,?,?,?);`

	getAllMahasiswa = `SELECT * FROM mahasiswa;`

	getMahasiswaByID = `SELECT * FROM mahasiswa WHERE nim = ?;`

	deleteMahasiswaByID = `DELETE FROM mahasiswa WHERE nim = ?;`

	updateMahasiswa = `UPDATE mahasiswa SET %s WHERE nim = ?;`
)

// NewSQLiteReadWriter is function creating connection with sql lite database.
// The request parameters for this function is db Host, db name, username, password.
// The sql connection was made by using sql.open() function.
// function returns database as response.
func NewSQLiteReadWriter() (_interface.ReadWriter, error) {
	db, err := sql.Open("sqlite3", "./mahasiswa.db")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	err = prepareDB(db)
	if err != nil {
		return nil, err
	}

	// db.SetMaxOpenConns(20)
	// db.SetMaxIdleConns(10)
	// db.SetConnMaxLifetime(30 * time.Minute)

	return &sqLiteReadWriter{sqLite: db}, nil
}

// Close is used for closing the sql connection
func (rw *sqLiteReadWriter) Close() error {
	if rw.sqLite != nil {
		if err := rw.sqLite.Close(); err != nil {
			return err
		}
		rw.sqLite = nil
	}

	return nil
}

func prepareDB(db *sql.DB) error {
	_, err := db.Exec(sqlTable)
	if err != nil {
		return err
	}

	return nil
}
