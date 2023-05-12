package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/debyltech/go-snipcart-api/config"
	"github.com/debyltech/go-snipcart/snipcart"
	"github.com/gin-gonic/gin"
)

func init() {
	var err error
	apiConfig, err = config.NewConfigFromEnv(true)
	if err != nil {
		fmt.Printf("ERROR %s\n", err.Error())
		return
	}

	// Initialize the snipcart client
	snipcartClient = snipcart.NewClient(apiConfig.SnipcartApiKey)

	// Set if we print logs for debugger.Print..
	debugger.Enabled = apiConfig.Production
	if apiConfig.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ready",
			"version": BuildVersion,
		})
	})
	api := r.Group("/api/snipcart")
	{
		api.GET("/products", GetProducts())
		api.GET("/products/:id", GetProductById())
	}

	ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
