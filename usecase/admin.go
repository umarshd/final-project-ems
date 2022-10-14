package usecase

import (
	"final-project-ems/entity"
	"final-project-ems/helpers"
	"final-project-ems/repository"
)

type IAdminUseCase interface {
	CreateAdmin(admin entity.RequestAdmin) (entity.Admin, error)
	LoginAdmin(admin entity.RequestLoginAdmin) (entity.ResponseAdminLogin, error)
}

type AdminUseCase struct {
	adminRepository repository.IAdminRepository
}

func NewAdminUseCase(adminRepository repository.IAdminRepository) *AdminUseCase {
	return &AdminUseCase{adminRepository}
}

func (useCase AdminUseCase) CreateAdmin(admin entity.RequestAdmin) (entity.ResponseAdmin, error) {
	passwordHash := helpers.HashPass(admin.Password)
	u := entity.Admin{
		Nama:     admin.Nama,
		Email:    admin.Email,
		Password: passwordHash,
	}

	adminResponse, err := useCase.adminRepository.Store(u)

	if err != nil {
		return entity.ResponseAdmin{}, err
	}

	userRes := entity.ResponseAdmin{
		UID:   adminResponse.UID,
		Nama:  adminResponse.Nama,
		Email: adminResponse.Email,
	}

	return userRes, nil

}

func (useCase AdminUseCase) LoginAdmin(admin entity.RequestLoginAdmin) (entity.ResponseAdminLogin, error) {
	adminResponse, err := useCase.adminRepository.Login(admin.Email)

	if err != nil {
		return entity.ResponseAdminLogin{}, err
	}

	checkPass := helpers.ComparePass([]byte(adminResponse.Password), []byte(admin.Password))

	if !checkPass {
		return entity.ResponseAdminLogin{}, err
	}

	token, _ := helpers.GenerateTokenAdmin(adminResponse.UID, adminResponse.Email)

	userRes := entity.ResponseAdminLogin{
		Token: token,
	}

	return userRes, nil
}
