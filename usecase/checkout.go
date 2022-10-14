package usecase

import (
	"final-project-ems/entity"
	"final-project-ems/repository"
)

type ICheckoutUseCase interface {
	StoreCheckout(checkout entity.Checkout) (entity.Checkout, error)
}

type CheckoutUseCase struct {
	checkoutRepository repository.ICheckoutRepository
}

func NewCheckoutUseCase(checkoutRepository repository.ICheckoutRepository) *CheckoutUseCase {
	return &CheckoutUseCase{checkoutRepository}
}

func (u CheckoutUseCase) StoreCheckout(checkoutRequest entity.Checkout) (entity.Checkout, error) {
	checkout := entity.Checkout{
		User_id:   checkoutRequest.User_id,
		Cart_id:   checkoutRequest.Cart_id,
		Image_url: checkoutRequest.Image_url,
		Total:     checkoutRequest.Total,
	}

	checkout, err := u.checkoutRepository.StoreCheckout(checkout)
	if err != nil {
		return checkout, err
	}

	return checkout, nil
}
