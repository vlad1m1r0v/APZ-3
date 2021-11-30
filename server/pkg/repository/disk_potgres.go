package repository

import (
	"database/sql"
)

type Disk struct {
	id int `json:"id" db:"id"`
	name string `json:"name" db:"name" binding:"required"`
	machinesId int `json:"machinesId" db:"machines_id"`
	capacity int `json:"capacity" db:"capacity" binding:"required"`
}

type diskPostgres struct {
	db *sql.DB
}

func newDiskPostgres(db *sql.DB) *diskPostgres{
	return &diskPostgres{db:db}
}

func (r *diskPostgres) ConnectToMachine(id, machineId int) error{
	_, err := r.db.Exec("UPDATE disks SET machines_id = $2 WHERE id = $1", id, machineId)
	return err
}

