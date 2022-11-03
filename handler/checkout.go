package handler

import (
	"final-project-ems/entity"
	"final-project-ems/usecase"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CheckoutHandler struct {
	checkoutUseCase *usecase.CheckoutUseCase
	cartUseCase     *usecase.CartUseCase
}

func NewCheckoutHandler(checkoutUseCase *usecase.CheckoutUseCase, cartUseCase *usecase.CartUseCase) *CheckoutHandler {
	return &CheckoutHandler{checkoutUseCase, cartUseCase}
}

// Chckout godoc
// @Summary Create Checkout
// @Description Create Checkout
// @Tags Checkout
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param body body entity.RequestCheckout true "Create Checkout"
// @Success 201 {object} entity.ResponseCheckout
// @Failure 400 {object} entity.ErrorResponse
// @Router /users/checkout [post]
func (handler CheckoutHandler) CreateCheckout(c echo.Context) error {
	req := entity.RequestCheckout{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	userId := c.Get("user").(jwt.MapClaims)["id"].(string)
	fmt.Println(userId, "ini")

	if userId != req.User_id {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "User not found",
		})
	}

	cart, err := handler.cartUseCase.FindCartById(req.Cart_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if cart.IsCheckout == true {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Cart has been checkout",
		})
	}

	if cart.User_id != req.User_id {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Cart not found",
		})
	}

	newReq := entity.Checkout{
		User_id: req.User_id,
		Cart_id: req.Cart_id,
		Total:   cart.Total,
	}

	checkout, err := handler.checkoutUseCase.StoreCheckout(newReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	reqCart := entity.Cart{
		ID:         cart.ID,
		User_id:    cart.User_id,
		Total:      cart.Total,
		IsCheckout: true,
	}

	updateCart, err := handler.cartUseCase.UpdateCart(reqCart)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_ = updateCart

	return c.JSON(http.StatusCreated, &entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Checkout successfully",
		Data: entity.ResponseCheckout{
			Id: checkout.Id,
		},
	})
}
