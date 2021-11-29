package repository

import "database/sql"

type Machine interface {

}

type Disk interface {

}

type Repository struct {
	Machine
	Disk
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
