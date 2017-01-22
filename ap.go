package main

import (
	"github.com/EncryptAny/AditumAllAPI/app"
	"github.com/goadesign/goa"
)

// ApController implements the ap resource.
type ApController struct {
	*goa.Controller
}

// NewApController creates a ap controller.
func NewApController(service *goa.Service) *ApController {
	return &ApController{Controller: service.NewController("ApController")}
}

// Create runs the create action.
func (c *ApController) Create(ctx *app.CreateApContext) error {
	// ApController_Create: start_implement

	// Put your logic here

	// ApController_Create: end_implement
	res := &app.Ap{}
	return ctx.OK(res)
}

// Downvote runs the downvote action.
func (c *ApController) Downvote(ctx *app.DownvoteApContext) error {
	// ApController_Downvote: start_implement

	// Put your logic here

	// ApController_Downvote: end_implement
	return nil
}

// Upvote runs the upvote action.
func (c *ApController) Upvote(ctx *app.UpvoteApContext) error {
	// ApController_Upvote: start_implement

	// Put your logic here

	// ApController_Upvote: end_implement
	return nil
}
