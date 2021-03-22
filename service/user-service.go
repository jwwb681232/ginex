package service

import (
	"GinRest/dto"
	"GinRest/entity"
	"GinRest/repository"
	"github.com/mashingan/smapping"
	"log"
)

type UserService interface {
	Update(user dto.UserUpdateDTO) entity.User
	Profile(userID string) entity.User
}

type userService struct {
	 userRepository repository.UserRepository
}

func (service userService) Update(user dto.UserUpdateDTO) entity.User {
	userToUpdate := entity.User{}
	err := smapping.FillStruct(&userToUpdate,smapping.MapFields(&user))
	if err != nil {
		log.Printf("Failed map %v:",err)
	}

	updateUser := service.userRepository.UpdateUser(userToUpdate)
	return updateUser
}

func (service userService) Profile(userID string) entity.User {
	return service.userRepository.ProfileUser(userID)
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}