package repository

import "database/sql"

type MachineRepository interface {
	GetAll()([]Machine, error)

}

type DiskRepository interface {
	ConnectToMachine(id, machineId int) error
}

type Repository struct {
	MachineRepository
	DiskRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		MachineRepository: newMachinePostgres(db),
		DiskRepository: newDiskPostgres(db),
	}
}
