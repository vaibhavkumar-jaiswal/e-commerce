// Package service provides the service layer for product management operations.
package service

import (
	"e-commerce/models"
	"e-commerce/modules/product/dtos"
	"e-commerce/modules/product/repository"
	"e-commerce/utils/helper"
	"fmt"

	"github.com/google/uuid"
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

// GetProductByID retrieves a product by their unique identifier (UUID).
//
// Parameters:
//
//	productID (string): The UUID of the product in string format.
//
// Returns:
//
//	any: Typically a product response object if found.
//	error: An error if the product is not found or if an error occurred during retrieval.
func (service *Service) GetProductByID(productID string) (any, error) {
	parsedUUID, err := uuid.Parse(productID)
	if err != nil {
		return nil, fmt.Errorf("invalid productID format, expects uuid")
	}

	product, err := service.productRepo.Get(parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return nil, fmt.Errorf(pgErr.Detail)
		}
		return nil, err
	}

	if product == nil {
		return nil, fmt.Errorf("no product found with productID = %s", parsedUUID)
	}

	return product.ResponseObj(), nil
}

// AddProduct adds a new product to the database.
// It takes a product request object, creates a new product model,
// and saves it to the database.
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

// UpdateProduct updates an existing product in the database.
func (service *Service) UpdateProduct(productID string, request dtos.UpdateProductRequest) (string, error) {
	parsedUUID, err := uuid.Parse(productID)
	if err != nil {
		return "", fmt.Errorf("invalid productID format, expects uuid")
	}

	product, err := service.productRepo.Get(parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	if product == nil {
		return "", fmt.Errorf("no product found with productID = %s", parsedUUID)
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.ProductCategoryID = request.ProductCategoryID
	product.Stock = request.Stock
	product.ImageURL = request.ImageURL

	err = service.productRepo.Update(product)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "Product upadated successfully.", nil
}

// PartialUpdateProduct updates specific fields of a product based on the provided patch request.
func (service *Service) PartialUpdateProduct(productID string, request dtos.PatchProductRequest) (string, error) {
	parsedUUID, err := uuid.Parse(productID)
	if err != nil {
		return "", fmt.Errorf("invalid productID format, expects uuid")
	}

	product, err := service.productRepo.Get(parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	if product == nil {
		return "", fmt.Errorf("no product found with productID = %s", parsedUUID)
	}

	patchData := make(map[string]any)

	if request.Name != nil {
		patchData["name"] = request.Name
	}

	if request.Description != nil {
		patchData["description"] = request.Description
	}

	if request.Price != nil {
		patchData["price"] = request.Price
	}

	if request.ProductCategoryID != nil {
		patchData["product_category_id"] = request.ProductCategoryID
	}

	if request.Stock != nil {
		patchData["stock"] = request.Stock
	}

	if request.ImageURL != nil {
		patchData["imageurl"] = request.ImageURL
	}

	err = service.productRepo.PartialUpdate(patchData, "products.product_id = ?", parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "Product upadated successfully.", nil
}

// DeleteProduct deletes a product from the database.
func (service *Service) DeleteProduct(productID string) (string, error) {
	parsedUUID, err := uuid.Parse(productID)
	if err != nil {
		return "", fmt.Errorf("invalid productID format, expects uuid")
	}

	product, err := service.productRepo.Get(parsedUUID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	if product == nil {
		return "", fmt.Errorf("no product found with productID = %s", parsedUUID)
	}

	err = service.productRepo.Delete(product, true)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			return "", fmt.Errorf(pgErr.Detail)
		}
		return "", err
	}

	return "Product successfully deleted.", nil
}
