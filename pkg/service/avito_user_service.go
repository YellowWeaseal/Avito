package service

import (
	"awesomeProject1"
	"awesomeProject1/pkg/repository"
)

type AvitoUserService struct {
	repo repository.AvitoUser
}

func NewAvitoUserService(repo repository.AvitoUser) *AvitoUserService {
	return &AvitoUserService{repo: repo}
}

func (s *AvitoUserService) Create(user awesomeProject1.UserUpdate) (int, error) {
	return s.repo.Create(user)
}

func (s *AvitoUserService) GetAll() ([]awesomeProject1.User, error) {
	return s.repo.GetAll()
}

func (s *AvitoUserService) GetById(userId int) (awesomeProject1.User, error) {
	return s.repo.GetById(userId)
}

func (s *AvitoUserService) Delete(userId int) error {
	return s.repo.Delete(userId)
}

func (s *AvitoUserService) Update(userId int, input awesomeProject1.UserUpdate) (int, error) {
	/*if err := input.Validate(); err != nil {
		return err
	}
	*/
	return s.repo.Update(userId, input)
}
