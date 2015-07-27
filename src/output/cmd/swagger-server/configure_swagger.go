package main

import (
	"github.com/go-swagger/go-swagger/errors"

	"output/models"
	"output/restapi/operations"
)

// This file is safe to edit. Once it exists it won't be overwritten

func configureAPI(api *operations.SwaggerAPI) {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = swagger.JSONConsumer()

	api.JSONProducer = swagger.JSONProducer()

	api.Handler = HandlerFunc(func() (*models.Model, error) {
		return nil, errors.NotImplemented("operation  has not yet been implemented")
	})

}
