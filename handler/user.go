package handler

import (
	"final-project-ems/entity"
	"final-project-ems/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

// Create User godoc
// @Summary Create User
// @Description Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body entity.RequestUser true "Create User"
// @Success 201 {object} entity.ResponseUser
// @Failure 400 {object} entity.ErrorResponse
// @Router /users/register [post]
func (h UserHandler) CreateUser(c echo.Context) error {
	req := entity.RequestUser{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := h.userUseCase.CreateUser(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "User created successfully",
		Data: entity.ResponseUser{
			UID:       user.UID,
			Name:      user.Name,
			Email:     user.Email,
			Gender:    user.Gender,
			Phone:     user.Phone,
			CreatedAt: time.Now(),
		},
	})
}

// Login User godoc
// @Summary Login User
// @Description Login User
// @Tags User
// @Accept  json
// @Produce  json
// @Param body body entity.RequestLoginUser true "Login User"
// @Success 200 {object} entity.ResponseUserLogin
// @Failure 400 {object} entity.ErrorResponse
// @Router /users/login [post]
func (h UserHandler) LoginUser(c echo.Context) error {
	req := entity.RequestLoginUser{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := h.userUseCase.LoginUser(req)

	if len(user.Token) < 1 {
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
		Message: "Login success",
		Data: entity.ResponseUserLogin{
			Token: user.Token,
		},
	})
}

// Create User Address godoc
// @Summary Create User Address
// @Description Create User Address
// @Tags User
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer Token"
// @Param path path string true "User ID"
// @Param body body entity.RequestUserAddress true "Create User Address"
// @Success 201 {object} entity.ResponseUserAddress
// @Failure 400 {object} entity.ErrorResponse
// @Router /users/{id}/address [post]

func (h UserHandler) CreateUserAddress(c echo.Context) error {
	req := entity.RequestUserAddress{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	user, err := h.userUseCase.CreateUserAddress(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "User address created successfully",
		Data: entity.ResponseUserAddress{
			ID:        user.ID,
			User_id:   user.User_id,
			Address:   user.Address,
			CreatedAt: time.Now(),
		},
	})
}
