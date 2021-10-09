// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// InventoryGetHandlerFunc turns a function with the right signature into a inventory get handler
type InventoryGetHandlerFunc func(InventoryGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn InventoryGetHandlerFunc) Handle(params InventoryGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// InventoryGetHandler interface for that can handle valid inventory get params
type InventoryGetHandler interface {
	Handle(InventoryGetParams, interface{}) middleware.Responder
}

// NewInventoryGet creates a new http.Handler for the inventory get operation
func NewInventoryGet(ctx *middleware.Context, handler InventoryGetHandler) *InventoryGet {
	return &InventoryGet{Context: ctx, Handler: handler}
}

/* InventoryGet swagger:route GET /store/inventory store inventoryGet

Returns pet inventories by status

*/
type InventoryGet struct {
	Context *middleware.Context
	Handler InventoryGetHandler
}

func (o *InventoryGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewInventoryGetParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
