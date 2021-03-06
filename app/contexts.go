// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/EncryptAny/AditumAllAPI/design
// --out=$(GOPATH)/src/github.com/EncryptAny/AditumAllAPI
// --version=v1.1.0-dirty
//
// API "aditum-all": Application Contexts
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateApContext provides the ap create action context.
type CreateApContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *NewAP
}

// NewCreateApContext parses the incoming request URL and body, performs validations and creates the
// context used by the ap controller create action.
func NewCreateApContext(ctx context.Context, service *goa.Service) (*CreateApContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateApContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CreateApContext) OK(r *Ap) error {
	ctx.ResponseData.Header().Set("Content-Type", "ap")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKVenue sends a HTTP response with status code 200.
func (ctx *CreateApContext) OKVenue(r *ApVenue) error {
	ctx.ResponseData.Header().Set("Content-Type", "ap")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *CreateApContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// DownvoteApContext provides the ap downvote action context.
type DownvoteApContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ApID int
}

// NewDownvoteApContext parses the incoming request URL and body, performs validations and creates the
// context used by the ap controller downvote action.
func NewDownvoteApContext(ctx context.Context, service *goa.Service) (*DownvoteApContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DownvoteApContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramApID := req.Params["apID"]
	if len(paramApID) > 0 {
		rawApID := paramApID[0]
		if apID, err2 := strconv.Atoi(rawApID); err2 == nil {
			rctx.ApID = apID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("apID", rawApID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DownvoteApContext) OK(r *Votes) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DownvoteApContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpvoteApContext provides the ap upvote action context.
type UpvoteApContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ApID int
}

// NewUpvoteApContext parses the incoming request URL and body, performs validations and creates the
// context used by the ap controller upvote action.
func NewUpvoteApContext(ctx context.Context, service *goa.Service) (*UpvoteApContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpvoteApContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramApID := req.Params["apID"]
	if len(paramApID) > 0 {
		rawApID := paramApID[0]
		if apID, err2 := strconv.Atoi(rawApID); err2 == nil {
			rctx.ApID = apID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("apID", rawApID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpvoteApContext) OK(r *Votes) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpvoteApContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// AllVenueContext provides the venue all action context.
type AllVenueContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewAllVenueContext parses the incoming request URL and body, performs validations and creates the
// context used by the venue controller all action.
func NewAllVenueContext(ctx context.Context, service *goa.Service) (*AllVenueContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := AllVenueContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *AllVenueContext) OK(r []*Venue) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// DownvoteAIVenueContext provides the venue downvoteAI action context.
type DownvoteAIVenueContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AiType  string
	VenueID int
}

// NewDownvoteAIVenueContext parses the incoming request URL and body, performs validations and creates the
// context used by the venue controller downvoteAI action.
func NewDownvoteAIVenueContext(ctx context.Context, service *goa.Service) (*DownvoteAIVenueContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := DownvoteAIVenueContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAiType := req.Params["aiType"]
	if len(paramAiType) > 0 {
		rawAiType := paramAiType[0]
		rctx.AiType = rawAiType
		if !(rctx.AiType == "epiPen" || rctx.AiType == "bathroom" || rctx.AiType == "aed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`aiType`, rctx.AiType, []interface{}{"epiPen", "bathroom", "aed"}))
		}
	}
	paramVenueID := req.Params["venueID"]
	if len(paramVenueID) > 0 {
		rawVenueID := paramVenueID[0]
		if venueID, err2 := strconv.Atoi(rawVenueID); err2 == nil {
			rctx.VenueID = venueID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("venueID", rawVenueID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *DownvoteAIVenueContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *DownvoteAIVenueContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// NewVenueContext provides the venue new action context.
type NewVenueContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *NewVenue
}

// NewNewVenueContext parses the incoming request URL and body, performs validations and creates the
// context used by the venue controller new action.
func NewNewVenueContext(ctx context.Context, service *goa.Service) (*NewVenueContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := NewVenueContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *NewVenueContext) OK(r *Venue) error {
	ctx.ResponseData.Header().Set("Content-Type", "venue")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKDetailed sends a HTTP response with status code 200.
func (ctx *NewVenueContext) OKDetailed(r *VenueDetailed) error {
	ctx.ResponseData.Header().Set("Content-Type", "venue")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// BadRequest sends a HTTP response with status code 400.
func (ctx *NewVenueContext) BadRequest() error {
	ctx.ResponseData.WriteHeader(400)
	return nil
}

// NewAIVenueContext provides the venue newAI action context.
type NewAIVenueContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AiType  string
	VenueID int
}

// NewNewAIVenueContext parses the incoming request URL and body, performs validations and creates the
// context used by the venue controller newAI action.
func NewNewAIVenueContext(ctx context.Context, service *goa.Service) (*NewAIVenueContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := NewAIVenueContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAiType := req.Params["aiType"]
	if len(paramAiType) > 0 {
		rawAiType := paramAiType[0]
		rctx.AiType = rawAiType
		if !(rctx.AiType == "epiPen" || rctx.AiType == "bathroom" || rctx.AiType == "aed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`aiType`, rctx.AiType, []interface{}{"epiPen", "bathroom", "aed"}))
		}
	}
	paramVenueID := req.Params["venueID"]
	if len(paramVenueID) > 0 {
		rawVenueID := paramVenueID[0]
		if venueID, err2 := strconv.Atoi(rawVenueID); err2 == nil {
			rctx.VenueID = venueID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("venueID", rawVenueID, "integer"))
		}
	}
	return &rctx, err
}

// Created sends a HTTP response with status code 201.
func (ctx *NewAIVenueContext) Created(r *Ai) error {
	ctx.ResponseData.Header().Set("Content-Type", "")
	return ctx.ResponseData.Service.Send(ctx.Context, 201, r)
}

// Conflict sends a HTTP response with status code 409.
func (ctx *NewAIVenueContext) Conflict() error {
	ctx.ResponseData.WriteHeader(409)
	return nil
}

// ShowVenueContext provides the venue show action context.
type ShowVenueContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	VenueID int
}

// NewShowVenueContext parses the incoming request URL and body, performs validations and creates the
// context used by the venue controller show action.
func NewShowVenueContext(ctx context.Context, service *goa.Service) (*ShowVenueContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowVenueContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramVenueID := req.Params["venueID"]
	if len(paramVenueID) > 0 {
		rawVenueID := paramVenueID[0]
		if venueID, err2 := strconv.Atoi(rawVenueID); err2 == nil {
			rctx.VenueID = venueID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("venueID", rawVenueID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowVenueContext) OK(r *Venue) error {
	ctx.ResponseData.Header().Set("Content-Type", "venue")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKDetailed sends a HTTP response with status code 200.
func (ctx *ShowVenueContext) OKDetailed(r *VenueDetailed) error {
	ctx.ResponseData.Header().Set("Content-Type", "venue")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowVenueContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// UpvoteAIVenueContext provides the venue upvoteAI action context.
type UpvoteAIVenueContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	AiType  string
	VenueID int
}

// NewUpvoteAIVenueContext parses the incoming request URL and body, performs validations and creates the
// context used by the venue controller upvoteAI action.
func NewUpvoteAIVenueContext(ctx context.Context, service *goa.Service) (*UpvoteAIVenueContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := UpvoteAIVenueContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramAiType := req.Params["aiType"]
	if len(paramAiType) > 0 {
		rawAiType := paramAiType[0]
		rctx.AiType = rawAiType
		if !(rctx.AiType == "epiPen" || rctx.AiType == "bathroom" || rctx.AiType == "aed") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`aiType`, rctx.AiType, []interface{}{"epiPen", "bathroom", "aed"}))
		}
	}
	paramVenueID := req.Params["venueID"]
	if len(paramVenueID) > 0 {
		rawVenueID := paramVenueID[0]
		if venueID, err2 := strconv.Atoi(rawVenueID); err2 == nil {
			rctx.VenueID = venueID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("venueID", rawVenueID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *UpvoteAIVenueContext) OK(resp []byte) error {
	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)
	_, err := ctx.ResponseData.Write(resp)
	return err
}

// NotFound sends a HTTP response with status code 404.
func (ctx *UpvoteAIVenueContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
