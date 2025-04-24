package socket

import (
	"main/infra"
)

type SocketServiceInterface interface{}

type SocketService struct {
	db infra.DBInterface
}

func NewSocketService(db infra.DBInterface) *SocketService {
	return &SocketService{
		db: db,
	}
}
