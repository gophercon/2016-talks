package main

import (
	"github.com/goadesign/goa"
	"github.com/raphael/gophercon2016/demos/1-design/app"
)

// BottleController implements the bottle resource.
type BottleController struct {
	*goa.Controller
}

// NewBottleController creates a bottle controller.
func NewBottleController(service *goa.Service) *BottleController {
	return &BottleController{Controller: service.NewController("BottleController")}
}

// Create runs the create action.
func (c *BottleController) Create(ctx *app.CreateBottleContext) error {
	// BottleController_Create: start_implement

	// Put your logic here

	// BottleController_Create: end_implement
	return nil
}

// Show runs the show action.
func (c *BottleController) Show(ctx *app.ShowBottleContext) error {
	// BottleController_Show: start_implement

	// Put your logic here

	// BottleController_Show: end_implement
	res := &app.Bottle{}
	return ctx.OK(res)
}
