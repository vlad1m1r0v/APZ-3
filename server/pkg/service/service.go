package service

type Machine interface {

}

type Disk interface {

}

type Service struct {

}

func NewService() *Service {
	return &Service{}
}
