package handler

import (
	"final-project-ems/entity"
	"final-project-ems/usecase"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ProductCategoryHandler struct {
	productCategoryUseCase *usecase.ProductCategoryUseCase
}

func NewProductCategoryHandler(productCategoryUseCase *usecase.ProductCategoryUseCase) *ProductCategoryHandler {
	return &ProductCategoryHandler{productCategoryUseCase}
}

// Create Product Category godoc
// @Summary Create Product Category
// @Description Create Product Category
// @Tags Product Category
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer Token"
// @Param body body entity.RequestProductCategory true "Create Product Category"
// @Success 201 {object} entity.ResponseProductCategory
// @Failure 400 {object} entity.ErrorResponse
// @Router /product-category [post]
func (h ProductCategoryHandler) CreateProductCategory(c echo.Context) error {
	req := entity.RequestProductCategory{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	productCategory, err := h.productCategoryUseCase.CreateProductCategory(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, &entity.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Product Category created successfully",
		Data: entity.ResponseProductCategory{
			ID:          productCategory.ID,
			Name:        productCategory.Name,
			Description: productCategory.Description,
			CreatedAt:   time.Now(),
		},
	})
}

// Get All Product Category godoc
// @Summary Get All Product Category
// @Description Get All Product Category
// @Tags Product Category
// @Accept  json
// @Produce  json
// @Param authorization header string true "Bearer Token"
// @Success 200 {object} entity.ResponseProductCategory
// @Failure 400 {object} entity.ErrorResponse
// @Router /product-category [get]
func (h ProductCategoryHandler) GetAllProductCategory(c echo.Context) error {
	productCategory, err := h.productCategoryUseCase.FindAllProductCategory()

	if err != nil {
		return c.JSON(http.StatusBadRequest, &entity.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if len(productCategory) < 0 {
		return c.JSON(http.StatusNotFound, &entity.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Product Category not found",
		})
	}

	return c.JSON(http.StatusOK, &entity.SuccessResponse{
		Status:  http.StatusOK,
		Message: "All Product Category",
		Data:    productCategory,
	})
}
