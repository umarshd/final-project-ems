package handler

import (
	"final-project-ems/entity"
	"final-project-ems/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	CartUseCase    *usecase.CartUseCase
	ProductUseCase *usecase.ProductUseCase
}

func NewCartHandler(cartUseCase *usecase.CartUseCase, productUseCase *usecase.ProductUseCase) *CartHandler {
	return &CartHandler{cartUseCase, productUseCase}
}

// Get All Cart By User godoc
// @Summary Get All Cart By User
// @Description Get All Cart By User
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} entity.Cart
// @Failure 400 {object} entity.ErrorResponse
// @Router /carts [get]
func (handler CartHandler) FindAllCart(c echo.Context) error {
	userId := c.Get("user").(jwt.MapClaims)["id"].(string)
	cartItem, err := handler.CartUseCase.FindAllCart(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entity.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Cart item found successfully",
		Data:    cartItem,
	})
}

// Create Cart godoc
// @Summary Create Cart
// @Description Create Cart
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Success 201 {object} entity.Cart
// @Failure 400 {object} entity.ErrorResponse
// @Router /carts [post]
func (handler CartHandler) CreateCart(c echo.Context) error {
	userId := c.Get("user").(jwt.MapClaims)["id"].(string)

	cart, err := handler.CartUseCase.CreateCart(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Cart created successfully",
		Data: entity.ResponseCart{
			ID:      cart.ID,
			User_id: cart.User_id,
			Total:   cart.Total,
			Created: time.Now(),
		},
	})
}

// Create Cart Item godoc
// @Summary Create Cart Item
// @Description Create Cart Item
// @Tags Cart
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param body body entity.RequestCartItem true "Create Cart Item"
// @Success 201 {object} entity.CartItem
// @Failure 400 {object} entity.ErrorResponse
// @Router /carts/items [post]
func (handler CartHandler) CreateCartItem(c echo.Context) error {
	req := entity.RequestCartItem{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	if err := c.Validate(req); err != nil {
		return err
	}

	checkCart, err := handler.CartUseCase.FindCartById(req.Cart_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userId := c.Get("user").(jwt.MapClaims)["id"].(string)

	if checkCart.User_id != userId {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Cart not found",
		})
	}

	if checkCart.IsCheckout == true {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Cart already checkout",
		})
	}

	cartItem, err := handler.CartUseCase.CreateCartItem(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})

	}

	product, err := handler.ProductUseCase.FindByIdProduct(cartItem.Product_id)
	cart, err := handler.CartUseCase.FindCartById(cartItem.Cart_id)

	priceProduct, _ := strconv.Atoi(product.Price)

	total := priceProduct * cartItem.Quantity

	reqUpdate := entity.Cart{
		ID:         cart.ID,
		User_id:    cart.User_id,
		Total:      cart.Total + total,
		IsCheckout: cart.IsCheckout,
	}

	newCart, err := handler.CartUseCase.UpdateCart(reqUpdate)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_ = newCart

	return c.JSON(http.StatusCreated, entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Cart item created successfully",
		Data: entity.ResponseCartItem{
			ID:         cartItem.ID,
			Cart_id:    cartItem.Cart_id,
			Product_id: cartItem.Product_id,
			Quantity:   cartItem.Quantity,
		},
	})
}
