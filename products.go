package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		products, err := snipcartClient.GetProducts(nil)
		if err != nil {
			debugger.Printf("issue with GetProducts: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, JSONErrorResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, products)
	}

	return fn
}

func GetProductById() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		product, err := snipcartClient.GetProductById(c.Param("id"))
		if err != nil {
			debugger.Printf("issue with GetProducts: %s\n", err.Error())
			c.JSON(http.StatusInternalServerError, JSONErrorResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, product)
	}

	return fn
}
