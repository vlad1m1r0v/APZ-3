package service

import "github.com/vlad1m1r0v/APZ-3/server/pkg/repository"

type machineService struct {
	repo repository.MachineRepository
}

func NewMachineService(repo repository.MachineRepository) *machineService {
	return &machineService{repo: repo}
}

func (s *machineService) GetAllMachines() ([]repository.Machine, error) {
	result , err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return result, nil
}