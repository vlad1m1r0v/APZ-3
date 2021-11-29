package service

import "github.com/vlad1m1r0v/APZ-3/server/pkg/repository"

type Machine interface {

}

type Disk interface {

}

type Service struct {

}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
