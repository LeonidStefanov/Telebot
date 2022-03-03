package service

import (
	"home/leonid/Git/Pract/telegram_bot/pkg/models"
	"log"
)

type Database interface {
	DeleteRequst(id int) error
	GetUserRequests(user_name string) ([]models.Request, error)
	AddRequest(userId int, userName, response, resTime string) error
	GetRequest() ([]models.Request, error)
	GetUsers() ([]models.Request, error)
	Close() error
}
type Service interface {
	DeleteRequst(id int) error
	GetUserRequests(user_name string) ([]models.Request, error)
	AddRequest(userId int, userName, response, resTime string) error
	GetRequest() ([]models.Request, error)
	GetUsers() ([]models.Request, error)
}

type service struct {
	db Database
}

func NewService(d Database) Service {
	return &service{
		db: d,
	}
}

func (s *service) AddRequest(userId int, userName, response, resTime string) error {

	err := s.db.AddRequest(userId, userName, response, resTime)
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

func (s *service) GetUserRequests(user_name string) ([]models.Request, error) {
	rep, err := s.db.GetUserRequests(user_name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rep, nil
}

func (s *service) GetUsers() ([]models.Request, error) {
	rep, err := s.db.GetRequest()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rep, nil
}

func (s *service) DeleteRequst(id int) error {
	err := s.db.DeleteRequst(id)

	if err != nil {
		return nil
	}

	return nil
}
