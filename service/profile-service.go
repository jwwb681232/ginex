package service

import (
	"GinRest/dto"
	"GinRest/entity"
	"GinRest/repository"
)

type ProfileService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(dto dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type profileService struct {
	userRepository repository.UserRepository
}

func NewProfileService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}