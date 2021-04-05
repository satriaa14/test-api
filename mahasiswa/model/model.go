package model

import "time"

type Mahasiswa struct {
	NIM       string    `json:"nim,omitempty"`
	Name      string    `json:"name,omitempty"`
	Class     string    `json:"class,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	CreatedBy string    `json:"created_by,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	UpdatedBy string    `json:"updated_by,omitempty"`
}
