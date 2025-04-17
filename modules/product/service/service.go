// Package service provides the service layer for product management operations.
package service

import (
	"e-commerce/models"
	"e-commerce/modules/product/dtos"
	"e-commerce/modules/product/repository"
	"e-commerce/utils/helper"
	"fmt"

	"github.com/lib/pq"
)

// Service defines the structure for the product service layer.
// It contains a repository instance to interact with the database.
type Service struct {
	productRepo         *repository.ProductRepo
	productCategoryRepo *repository.ProductCategoryRepo
}

// NewUserService creates and returns a new User Service instance by initializing the repository.
// Returns:
//
//	*Service: A pointer to a new Service instance with its repository initialized.
func NewUserService() *Service {
	return &Service{
		productRepo:         repository.NewProductRepository(),
		productCategoryRepo: repository.NewProductCategoryRepository(),
	}
}

// GetProducts retrieves a list of products matching the specified query parameters.
// It builds a dynamic query, and returns a formatted list of product responses.
//
// Parameters:
//
//	queryParams (*dtos.ProductQueryParams): The query parameters for filtering products.
//
// Returns:
//
//	[]models.ProductResponse: A slice of product response objects.
//	error: An error if no data is found or if any operation fails.
func (service *Service) GetProducts(queryParams *dtos.ProductQueryParams) ([]models.ProductResponse, error) {

	filter := service.productRepo.GetFilter()

	filter = helper.BuildQuery(filter, queryParams)

	products, _, err := service.productRepo.FindAll(filter, "name", 0, 0)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if len(products) < 1 {
		return nil, fmt.Errorf("no data found")
	}

	return models.ProductList(products).ResponseList(), nil
}

func (service *Service) AddProduct(productRequest *dtos.ProductRequest) (string, error) {

	product := &models.Product{
		Name:              productRequest.Name,
		Description:       productRequest.Description,
		Price:             productRequest.Price,
		ProductCategoryID: productRequest.ProductCategoryID,
		Stock:             productRequest.Stock,
		ImageURL:          productRequest.ImageURL,
	}

	err := service.productRepo.Create(product)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "Product added successfully.", nil
}
