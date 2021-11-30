package repository

import (
	"database/sql"
)

type Machine struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	CpuCount string `json:"cpuCount" db:"cpucount"`
	TotalDiskSpace string `json:"totalDiskSpace" db:"totaldiskspace"`
}

type machinePostgres struct {
	db *sql.DB
}

func newMachinePostgres(db *sql.DB) *machinePostgres{
	return &machinePostgres{db:db}
}

func (r *machinePostgres) GetAll()([]Machine, error){
	rows, err := r.db.Query("select * from machines")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var machines [] Machine
	for rows.Next() {
		var m Machine
		if err := rows.Scan(&m.Id, &m.Name, &m.CpuCount, &m.TotalDiskSpace); err != nil {
			return nil, err
		}
		machines = append(machines, m)
	}
	if machines == nil {
		machines = make([]Machine, 0)
	}
	return machines, nil
}