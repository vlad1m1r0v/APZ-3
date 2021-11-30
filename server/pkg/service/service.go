package service

import "github.com/vlad1m1r0v/APZ-3/server/pkg/repository"

type MachineService interface {
	GetAllMachines() ([]repository.Machine, error)
}

type DiskService interface {
	ConnectToMachine(id, machineId int) error
}

type Service struct {
	MachineService
	DiskService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		MachineService: NewMachineService(repos.MachineRepository),
		DiskService: NewDiskService(repos.DiskRepository),
	}
}




