package usecase

import (
	"final-project-ems/entity"
	"final-project-ems/helpers"
	"final-project-ems/repository"
	"time"
)

type IUserUseCase interface {
	CreateUser(user entity.RequestUser) (entity.User, error)
	LoginUser(user entity.RequestLoginUser) (entity.ResponseUserLogin, error)
	CreateUserAddress(address entity.RequestUserAddress) (entity.UserAddress, error)
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{userRepository}
}

func (useCase UserUseCase) CreateUser(user entity.RequestUser) (entity.ResponseUser, error) {
	uProfile := entity.UserProfile{
		Name:   user.Name,
		Gender: user.Gender,
		Phone:  user.Phone,
	}

	passwordHash := helpers.HashPass(user.Password)

	userResponse, err := useCase.userRepository.StoreProfile(uProfile)

	if err != nil {
		return entity.ResponseUser{}, err
	}

	u := entity.User{
		Email:      user.Email,
		Password:   passwordHash,
		Profile_id: userResponse.ID,
	}

	users, err := useCase.userRepository.Store(u)

	if err != nil {
		return entity.ResponseUser{}, err
	}

	userRes := entity.ResponseUser{
		UID:       users.UID,
		Email:     users.Email,
		Name:      userResponse.Name,
		Gender:    userResponse.Gender,
		Phone:     userResponse.Phone,
		CreatedAt: time.Now(),
	}

	return userRes, nil
}

func (useCase UserUseCase) LoginUser(user entity.RequestLoginUser) (entity.ResponseUserLogin, error) {
	userResponse, err := useCase.userRepository.Login(user.Email)

	if err != nil {
		return entity.ResponseUserLogin{}, err
	}

	checkPass := helpers.ComparePass([]byte(userResponse.Password), []byte(user.Password))

	if !checkPass {
		return entity.ResponseUserLogin{}, err
	}

	token, _ := helpers.GenerateTokenUser(userResponse.UID, userResponse.Email)

	userRes := entity.ResponseUserLogin{
		Token: token,
	}

	return userRes, nil
}

func (useCase UserUseCase) CreateUserAddress(address entity.RequestUserAddress) (entity.UserAddress, error) {
	userAddress := entity.UserAddress{
		User_id: address.User_id,
		Address: address.Address,
	}

	userAddressResponse, err := useCase.userRepository.StoreAddress(userAddress)

	if err != nil {
		return entity.UserAddress{}, err
	}

	return userAddressResponse, nil
}
