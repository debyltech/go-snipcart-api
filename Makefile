SRC=$(wildcard *.go)

go-snipcart-api: $(SRC)
	CGO_ENABLED=0 go build -tags netgo

go-snipcart-api.zip: go-snipcart-api
	zip go-snipcart-api.zip go-snipcart-api

deploy-dev: go-snipcart-api.zip
	aws --profile debyltech lambda update-function-code --function-name 'snipcart-api-dev' --zip-file 'fileb://go-snipcart-api.zip'

deploy-prod: go-snipcart-api.zip
	aws --profile debyltech lambda update-function-code --function-name 'snipcart-api-prod' --zip-file 'fileb://go-snipcart-api.zip'