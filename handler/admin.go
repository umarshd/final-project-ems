package handler

import (
	"final-project-ems/entity"
	"final-project-ems/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminUseCase *usecase.AdminUseCase
}

func NewAdminHandler(adminUseCase *usecase.AdminUseCase) *AdminHandler {
	return &AdminHandler{adminUseCase: adminUseCase}
}

// Create Admin godoc
// @Summary Create Admin
// @Description Create Admin
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param body body entity.RequestAdmin true "Create Admin"
// @Success 201 {object} entity.ResponseAdmin
// @Failure 400 {object} entity.ErrorResponse
// @Router /admin/register [post]
func (handler AdminHandler) CreateAdmin(c echo.Context) error {
	req := entity.RequestAdmin{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := handler.adminUseCase.CreateAdmin(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Data: entity.ResponseAdmin{
			UID:       user.UID,
			Nama:      user.Nama,
			Email:     user.Email,
			CreatedAt: time.Now(),
		},
	})
}

// Login Admin godoc
// @Summary Login Admin
// @Description Login Admin
// @Tags Admin
// @Accept  json
// @Produce  json
// @Param body body entity.RequestLoginAdmin true "Login Admin"
// @Success 200 {object} entity.ResponseAdminLogin
// @Failure 400 {object} entity.ErrorResponse
// @Router /admin/login [post]
func (handler AdminHandler) LoginAdmin(c echo.Context) error {
	req := entity.RequestLoginAdmin{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	admin, err := handler.adminUseCase.LoginAdmin(req)

	if len(admin.Token) < 1 {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Email or password is wrong",
		})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, &entity.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Login successfully",
		Data: entity.ResponseAdminLogin{
			Token: admin.Token,
		},
	})
}
