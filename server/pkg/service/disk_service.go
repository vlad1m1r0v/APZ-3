package service

import "github.com/vlad1m1r0v/APZ-3/server/pkg/repository"

type diskService struct {
	repo repository.DiskRepository

}

func NewDiskService(repo repository.DiskRepository) *diskService {
	return &diskService{repo: repo}
}

func (s *diskService) ConnectToMachine(id, machineId int) error {
	err := s.repo.ConnectToMachine(id, machineId)
	if err != nil {
		return err
	}
	return nil
}
