package main

import (
	"context"

	"github.com/merynayr/mall-tenants/internal/app"
	"github.com/merynayr/mall-tenants/internal/logger"
)

// @title mall-tenants
// @version 1.0
// @description This is a mall-tenants API
// @contact.name Dmitry Boyarkin
// @contact.email boyarkin_dima2@mail.ru
// @host localhost:8080
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		logger.Error("Error", "failed to init app: ", err.Error())
	}

	err = a.Run()
	if err != nil {
		logger.Error("Error:", "failed to run app:", err.Error())
	}
}
