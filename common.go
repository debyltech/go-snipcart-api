package main

import (
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/debyltech/go-snipcart-api/config"
	"github.com/debyltech/go-snipcart-api/debug"
	"github.com/debyltech/go-snipcart/snipcart"
)

const (
	UriSnipcartApi string = "https://app.snipcart.com/api/"
	UriProductsApi        = UriSnipcartApi + "products"
)

var (
	apiConfig      *config.Config
	debugger       debug.Debug
	ginLambda      *ginadapter.GinLambda
	snipcartClient *snipcart.Client

	BuildVersion string
)

type JSONErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
