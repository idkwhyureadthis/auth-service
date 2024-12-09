package service

import database "github.com/idkwhyureadthis/auth-service/pkg/database"

type Service struct {
	db     *database.DB
	secret []byte
}

func New(connUrl string, secret []byte) *Service {
	s := &Service{
		db:     database.New(connUrl),
		secret: secret,
	}
	return s
}
