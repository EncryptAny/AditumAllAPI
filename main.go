//go:generate goagen bootstrap -d github.com/EncryptAny/AditumAllAPI/design

package main

import (
	"github.com/EncryptAny/AditumAllAPI/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("aditum-all")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "ap" controller
	c := NewApController(service)
	app.MountApController(service, c)
	// Mount "venue" controller
	c2 := NewVenueController(service)
	app.MountVenueController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
