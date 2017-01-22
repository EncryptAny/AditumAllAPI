package main

import (
	"github.com/EncryptAny/AditumAllAPI/app"
	"github.com/goadesign/goa"
)

// VenueController implements the venue resource.
type VenueController struct {
	*goa.Controller
}

// NewVenueController creates a venue controller.
func NewVenueController(service *goa.Service) *VenueController {
	return &VenueController{Controller: service.NewController("VenueController")}
}

// All runs the all action.
func (c *VenueController) All(ctx *app.AllVenueContext) error {
	// VenueController_All: start_implement

	// Put your logic here

	// VenueController_All: end_implement
	return nil
}

// DownvoteAI runs the downvoteAI action.
func (c *VenueController) DownvoteAI(ctx *app.DownvoteAIVenueContext) error {
	// VenueController_DownvoteAI: start_implement

	// Put your logic here

	// VenueController_DownvoteAI: end_implement
	return nil
}

// New runs the new action.
func (c *VenueController) New(ctx *app.NewVenueContext) error {
	// VenueController_New: start_implement

	// Put your logic here

	// VenueController_New: end_implement
	res := &app.Venue{}
	return ctx.OK(res)
}

// NewAI runs the newAI action.
func (c *VenueController) NewAI(ctx *app.NewAIVenueContext) error {
	// VenueController_NewAI: start_implement

	// Put your logic here

	// VenueController_NewAI: end_implement
	return nil
}

// Show runs the show action.
func (c *VenueController) Show(ctx *app.ShowVenueContext) error {
	// VenueController_Show: start_implement

	// Put your logic here

	// VenueController_Show: end_implement
	res := &app.Venue{}
	return ctx.OK(res)
}

// UpvoteAI runs the upvoteAI action.
func (c *VenueController) UpvoteAI(ctx *app.UpvoteAIVenueContext) error {
	// VenueController_UpvoteAI: start_implement

	// Put your logic here

	// VenueController_UpvoteAI: end_implement
	return nil
}
