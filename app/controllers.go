// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/EncryptAny/AditumAllAPI/design
// --out=$(GOPATH)/src/github.com/EncryptAny/AditumAllAPI
// --version=v1.1.0-dirty
//
// API "aditum-all": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ApController is the controller interface for the Ap actions.
type ApController interface {
	goa.Muxer
	Create(*CreateApContext) error
	Downvote(*DownvoteApContext) error
	Upvote(*UpvoteApContext) error
}

// MountApController "mounts" a Ap resource controller on the given service.
func MountApController(service *goa.Service, ctrl ApController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateApContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*NewAP)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/ap", ctrl.MuxHandler("Create", h, unmarshalCreateApPayload))
	service.LogInfo("mount", "ctrl", "Ap", "action", "Create", "route", "POST /ap")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDownvoteApContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Downvote(rctx)
	}
	service.Mux.Handle("POST", "/ap/:apID/upvote", ctrl.MuxHandler("Downvote", h, nil))
	service.LogInfo("mount", "ctrl", "Ap", "action", "Downvote", "route", "POST /ap/:apID/upvote")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpvoteApContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Upvote(rctx)
	}
	service.Mux.Handle("POST", "/ap/:apID/downvote", ctrl.MuxHandler("Upvote", h, nil))
	service.LogInfo("mount", "ctrl", "Ap", "action", "Upvote", "route", "POST /ap/:apID/downvote")
}

// unmarshalCreateApPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateApPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &newAP{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// VenueController is the controller interface for the Venue actions.
type VenueController interface {
	goa.Muxer
	All(*AllVenueContext) error
	DownvoteAI(*DownvoteAIVenueContext) error
	New(*NewVenueContext) error
	NewAI(*NewAIVenueContext) error
	Show(*ShowVenueContext) error
	UpvoteAI(*UpvoteAIVenueContext) error
}

// MountVenueController "mounts" a Venue resource controller on the given service.
func MountVenueController(service *goa.Service, ctrl VenueController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewAllVenueContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.All(rctx)
	}
	service.Mux.Handle("GET", "/venue", ctrl.MuxHandler("All", h, nil))
	service.LogInfo("mount", "ctrl", "Venue", "action", "All", "route", "GET /venue")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDownvoteAIVenueContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.DownvoteAI(rctx)
	}
	service.Mux.Handle("POST", "/venue/:venueID/:aiType/downvote", ctrl.MuxHandler("DownvoteAI", h, nil))
	service.LogInfo("mount", "ctrl", "Venue", "action", "DownvoteAI", "route", "POST /venue/:venueID/:aiType/downvote")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewNewVenueContext(ctx, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*NewVenue)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.New(rctx)
	}
	service.Mux.Handle("POST", "/venue", ctrl.MuxHandler("New", h, unmarshalNewVenuePayload))
	service.LogInfo("mount", "ctrl", "Venue", "action", "New", "route", "POST /venue")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewNewAIVenueContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.NewAI(rctx)
	}
	service.Mux.Handle("POST", "/venue/:venueID/:aiType", ctrl.MuxHandler("NewAI", h, nil))
	service.LogInfo("mount", "ctrl", "Venue", "action", "NewAI", "route", "POST /venue/:venueID/:aiType")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowVenueContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/venue/:venueID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Venue", "action", "Show", "route", "GET /venue/:venueID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpvoteAIVenueContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.UpvoteAI(rctx)
	}
	service.Mux.Handle("POST", "/venue/:venueID/:aiType/upvote", ctrl.MuxHandler("UpvoteAI", h, nil))
	service.LogInfo("mount", "ctrl", "Venue", "action", "UpvoteAI", "route", "POST /venue/:venueID/:aiType/upvote")
}

// unmarshalNewVenuePayload unmarshals the request body into the context request data Payload field.
func unmarshalNewVenuePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &newVenue{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}