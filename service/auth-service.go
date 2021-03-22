package service

import (
	"GinRest/dto"
	"GinRest/entity"
	"GinRest/repository"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService interface {
	VerifyCredential(email, password string) interface{}
	CreateUser(dto dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type authService struct {
	userRepository repository.UserRepository
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (service *authService) VerifyCredential(email, password string) interface{} {
	user := service.userRepository.FindByEmail(email)
	if comparedPassword(user.Password,[]byte(password)) {
		return user
	}
	return false
}

/*func (service *authService) VerifyCredential(email, password string) interface{} {
	res := service.userRepository.VerifyCredential(email,password)
	if v, ok := res.(entity.User); ok {
		comparedPassword := comparedPassword(v.Password,[]byte(password))
		if v.Email == email && comparedPassword{
			return res
		}
		return false
	}
	return false
}*/

func (service *authService) CreateUser(dto dto.RegisterDTO) entity.User {
	userToCreate := entity.User{}
	err := smapping.FillStruct(&userToCreate,smapping.MapFields(&dto))
	if err != nil {
		log.Fatalf("Failed map %v",err)
	}
	res := service.userRepository.InsertUser(userToCreate)
	return res
}
func (service *authService) FindByEmail(email string) entity.User {
	return service.userRepository.FindByEmail(email)
}

func (service *authService) IsDuplicateEmail(email string) bool {
	res := service.userRepository.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparedPassword(hashed string,plainPassword []byte) bool {
	byteHash := []byte(hashed)
	err := bcrypt.CompareHashAndPassword(byteHash,plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}