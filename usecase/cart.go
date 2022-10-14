package usecase

import (
	"final-project-ems/entity"
	"final-project-ems/repository"
)

type ICartUseCase interface {
	FindAllCart(userId string) ([]entity.Cart, error)
	CreateCart(userId string) (entity.Cart, error)
	FindById(id int) (entity.Cart, error)
	UpdateCart(cart entity.Cart) (entity.Cart, error)
	CreateCartItem(cartItem entity.RequestCartItem) (entity.CartItem, error)
}

type CartUseCase struct {
	cartRepository repository.ICartRepository
}

func NewCartUseCase(cartRepository repository.ICartRepository) *CartUseCase {
	return &CartUseCase{cartRepository}
}

func (u CartUseCase) FindAllCart(userId string) ([]entity.Cart, error) {
	carts, err := u.cartRepository.FindAllCart(userId)
	if err != nil {
		return carts, err
	}

	return carts, nil
}

func (u CartUseCase) CreateCart(userId string) (entity.Cart, error) {
	cart := entity.Cart{
		User_id:    userId,
		IsCheckout: false,
		Total:      0,
	}
	cart, err := u.cartRepository.StoreCart(cart)
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (u CartUseCase) FindCartById(id int) (entity.Cart, error) {
	cart, err := u.cartRepository.FindCartById(id)
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (u CartUseCase) UpdateCart(cartRequest entity.Cart) (entity.Cart, error) {

	cart, err := u.cartRepository.UpdateCart(cartRequest)
	if err != nil {
		return cart, err
	}

	return cart, nil
}

func (u CartUseCase) CreateCartItem(cartReques entity.RequestCartItem) (entity.CartItem, error) {

	newCartItem := entity.CartItem{
		Cart_id:    cartReques.Cart_id,
		Product_id: cartReques.Product_id,
		Quantity:   cartReques.Quantity,
	}
	cartItem, err := u.cartRepository.StoreCartItem(newCartItem)
	if err != nil {
		return cartItem, err
	}

	return cartItem, nil
}
