package service

import (
	"home/leonid/Git/Pract/telegram_bot/pkg/models"
	"log"
)

type Database interface {
	AddRequest(userId int, userName, response string) error
	GetRequest() ([]models.Request, error)
	Close() error
}
type Service interface {
	AddRequest(userId int, userName, response string) error
	GetRequest() ([]models.Request, error)
}

type service struct {
	db Database
}

func NewService(d Database) Service {
	return &service{
		db: d,
	}
}

func (s *service) AddRequest(userId int, userName, response string) error {
	err := s.db.AddRequest(userId, userName, response)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *service) GetRequest() ([]models.Request, error) {
	rep, err := s.db.GetRequest()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rep, nil
}
