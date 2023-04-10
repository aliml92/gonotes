package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


var products = []Product{
	{ID: "1", Name: "Apple", Price: 4.99},
	{ID: "2", Name: "Banana", Price: 2.99},
	{ID: "3", Name: "Tomato", Price: 3.99},
	{ID: "4", Name: "Kiwi", Price: 4.99},
	{ID: "5", Name: "Pineapple", Price: 3.99},
}


type Product struct {
	ID 			string    `json:"id"`
	Name 		string    `json:"name"`
	Price 		float32   `json:"price"`
}

type updateProductRequest struct {
	Product struct {
		Price float32 `json:"price"`
	} `json:"product"`
}

type updateProductResponse struct {
	Product Product `json:"product"`
}

func newProductResponse(product Product) updateProductResponse {
	return updateProductResponse{
		Product: product,
	}
}

func updateProductInStore(productID string, req updateProductRequest) (Product, error) {
	for i, product := range products {
		if product.ID == productID {
			product.Price = req.Product.Price
			products[i] = product
			return product, nil
		}
	}
	return Product{}, fmt.Errorf("product not found")
}

func UpdateProduct(c *gin.Context) {
	productID := c.Param("product_id")
	var req updateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("req: %+v", req)
	updatedProduct, err := updateProductInStore(productID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newProductResponse(updatedProduct))
}


func setupRouter() *gin.Engine {
	r := gin.Default()
	r.PATCH("/api/v1/products/:product_id", UpdateProduct)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}

