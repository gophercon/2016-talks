//************************************************************************//
// API "cellar": Application Contexts
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/raphael/gophercon2016/demos/2-implement/design
// --out=$(GOPATH)/src/github.com/raphael/gophercon2016/demos/2-implement
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"strconv"
)

// CreateBottleContext provides the bottle create action context.
type CreateBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Payload *CreateBottlePayload
}

// NewCreateBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller create action.
func NewCreateBottleContext(ctx context.Context, service *goa.Service) (*CreateBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := CreateBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// createBottlePayload is the bottle create action payload.
type createBottlePayload struct {
	// Name of bottle
	Name *string `json:"name,omitempty" xml:"name,omitempty" form:"name,omitempty"`
	// Rating of bottle
	Rating *int `json:"rating,omitempty" xml:"rating,omitempty" form:"rating,omitempty"`
	// Vintage of bottle
	Vintage *int `json:"vintage,omitempty" xml:"vintage,omitempty" form:"vintage,omitempty"`
}

// Validate runs the validation rules defined in the design.
func (payload *createBottlePayload) Validate() (err error) {
	if payload.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`payload`, "name"))
	}
	if payload.Vintage == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`payload`, "vintage"))
	}
	if payload.Rating == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`payload`, "rating"))
	}

	if payload.Name != nil {
		if len(*payload.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`payload.name`, *payload.Name, len(*payload.Name), 1, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`payload.rating`, *payload.Rating, 1, true))
		}
	}
	if payload.Rating != nil {
		if *payload.Rating > 5 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`payload.rating`, *payload.Rating, 5, false))
		}
	}
	if payload.Vintage != nil {
		if *payload.Vintage < 1900 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`payload.vintage`, *payload.Vintage, 1900, true))
		}
	}
	return
}

// Publicize creates CreateBottlePayload from createBottlePayload
func (payload *createBottlePayload) Publicize() *CreateBottlePayload {
	var pub CreateBottlePayload
	if payload.Name != nil {
		pub.Name = *payload.Name
	}
	if payload.Rating != nil {
		pub.Rating = *payload.Rating
	}
	if payload.Vintage != nil {
		pub.Vintage = *payload.Vintage
	}
	return &pub
}

// CreateBottlePayload is the bottle create action payload.
type CreateBottlePayload struct {
	// Name of bottle
	Name string `json:"name" xml:"name" form:"name"`
	// Rating of bottle
	Rating int `json:"rating" xml:"rating" form:"rating"`
	// Vintage of bottle
	Vintage int `json:"vintage" xml:"vintage" form:"vintage"`
}

// Validate runs the validation rules defined in the design.
func (payload *CreateBottlePayload) Validate() (err error) {
	if payload.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`payload`, "name"))
	}

	if len(payload.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`payload.name`, payload.Name, len(payload.Name), 1, true))
	}
	if payload.Rating < 1 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`payload.rating`, payload.Rating, 1, true))
	}
	if payload.Rating > 5 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`payload.rating`, payload.Rating, 5, false))
	}
	if payload.Vintage < 1900 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`payload.vintage`, payload.Vintage, 1900, true))
	}
	return
}

// Created sends a HTTP response with status code 201.
func (ctx *CreateBottleContext) Created() error {
	ctx.ResponseData.WriteHeader(201)
	return nil
}

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	ID int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(ctx context.Context, service *goa.Service) (*ShowBottleContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	rctx := ShowBottleContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramID := req.Params["id"]
	if len(paramID) > 0 {
		rawID := paramID[0]
		if id, err2 := strconv.Atoi(rawID); err2 == nil {
			rctx.ID = id
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("id", rawID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *Bottle) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gophercon.goa.bottle")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}
